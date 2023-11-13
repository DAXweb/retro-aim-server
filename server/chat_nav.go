package server

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mkaminski/goaim/oscar"
	"github.com/mkaminski/goaim/user"
	"io"
	"log/slog"
	"time"
)

type ChatNavHandler interface {
	CreateRoomHandler(ctx context.Context, sess *user.Session, newChatRoom ChatRoomFactory, snacPayloadIn oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.XMessage, error)
	RequestChatRightsHandler(ctx context.Context) oscar.XMessage
	RequestRoomInfoHandler(ctx context.Context, snacPayloadIn oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.XMessage, error)
}

func NewChatNavRouter(logger *slog.Logger, cr *ChatRegistry) ChatNavRouter {
	return ChatNavRouter{
		ChatNavHandler: ChatNavService{
			Logger: logger,
			cr:     cr,
		},
		RouteLogger: RouteLogger{
			Logger: logger,
		},
	}
}

type ChatNavRouter struct {
	ChatNavHandler
	RouteLogger
}

func (rt *ChatNavRouter) RouteChatNav(ctx context.Context, sess *user.Session, SNACFrame oscar.SnacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	switch SNACFrame.SubGroup {
	case oscar.ChatNavRequestChatRights:
		outSNAC := rt.RequestChatRightsHandler(ctx)
		rt.logRequestAndResponse(ctx, SNACFrame, nil, outSNAC.SnacFrame, outSNAC.SnacOut)
		return writeOutSNAC(SNACFrame, outSNAC.SnacFrame, outSNAC.SnacOut, sequence, w)
	case oscar.ChatNavRequestRoomInfo:
		inSNAC := oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo{}
		if err := oscar.Unmarshal(&inSNAC, r); err != nil {
			return err
		}
		outSNAC, err := rt.RequestRoomInfoHandler(ctx, inSNAC)
		if err != nil {
			return err
		}
		rt.logRequestAndResponse(ctx, SNACFrame, inSNAC, outSNAC.SnacFrame, outSNAC.SnacOut)
		return writeOutSNAC(SNACFrame, outSNAC.SnacFrame, outSNAC.SnacOut, sequence, w)
	case oscar.ChatNavCreateRoom:
		inSNAC := oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{}
		if err := oscar.Unmarshal(&inSNAC, r); err != nil {
			return err
		}
		outSNAC, err := rt.CreateRoomHandler(ctx, sess, NewChatRoom, inSNAC)
		if err != nil {
			return err
		}
		roomName, _ := inSNAC.GetString(oscar.ChatTLVRoomName)
		rt.Logger.InfoContext(ctx, "user started a chat room", slog.String("roomName", roomName))
		rt.logRequestAndResponse(ctx, SNACFrame, inSNAC, outSNAC.SnacFrame, outSNAC.SnacOut)
		return writeOutSNAC(SNACFrame, outSNAC.SnacFrame, outSNAC.SnacOut, sequence, w)
	default:
		return ErrUnsupportedSubGroup
	}
}

type ChatNavService struct {
	Logger *slog.Logger
	cr     *ChatRegistry
}

type ChatCookie struct {
	Cookie []byte `len_prefix:"uint16"`
	SessID string `len_prefix:"uint16"`
}

func (s ChatNavService) RequestChatRightsHandler(context.Context) oscar.XMessage {
	return oscar.XMessage{
		SnacFrame: oscar.SnacFrame{
			FoodGroup: oscar.CHAT_NAV,
			SubGroup:  oscar.ChatNavNavInfo,
		},
		SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(0x02, uint8(10)),
					oscar.NewTLV(0x03, oscar.SNAC_0x0D_0x09_TLVExchangeInfo{
						Identifier: 4,
						TLVBlock: oscar.TLVBlock{
							TLVList: oscar.TLVList{
								oscar.NewTLV(0x0002, uint16(0x0010)),
								oscar.NewTLV(0x00c9, uint16(15)),
								oscar.NewTLV(0x00d3, "default Exchange"),
								oscar.NewTLV(0x00d5, uint8(2)),
								oscar.NewTLV(0xd6, "us-ascii"),
								oscar.NewTLV(0xd7, "en"),
								oscar.NewTLV(0xd8, "us-ascii"),
								oscar.NewTLV(0xd9, "en"),
							},
						},
					}),
				},
			},
		},
	}
}

func NewChatRoom(logger *slog.Logger) ChatRoom {
	return ChatRoom{
		Cookie:         uuid.New().String(),
		CreateTime:     time.Now(),
		SessionManager: NewSessionManager(logger),
	}
}

type ChatRoomFactory func(logger *slog.Logger) ChatRoom

func (s ChatNavService) CreateRoomHandler(_ context.Context, sess *user.Session, newChatRoom ChatRoomFactory, snacPayloadIn oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.XMessage, error) {
	name, hasName := snacPayloadIn.GetString(oscar.ChatTLVRoomName)
	if !hasName {
		return oscar.XMessage{}, errors.New("unable to find chat name")
	}

	room := newChatRoom(s.Logger)
	room.DetailLevel = snacPayloadIn.DetailLevel
	room.Exchange = snacPayloadIn.Exchange
	room.InstanceNumber = snacPayloadIn.InstanceNumber
	room.Name = name
	s.cr.Register(room)

	// add user to chat room
	room.NewSessionWithSN(sess.ID(), sess.ScreenName())

	return oscar.XMessage{
		SnacFrame: oscar.SnacFrame{
			FoodGroup: oscar.CHAT_NAV,
			SubGroup:  oscar.ChatNavNavInfo,
		},
		SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(oscar.ChatNavTLVRoomInfo, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       snacPayloadIn.Exchange,
						Cookie:         room.Cookie,
						InstanceNumber: snacPayloadIn.InstanceNumber,
						DetailLevel:    snacPayloadIn.DetailLevel,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					}),
				},
			},
		},
	}, nil
}

func (s ChatNavService) RequestRoomInfoHandler(_ context.Context, snacPayloadIn oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.XMessage, error) {
	room, err := s.cr.Retrieve(string(snacPayloadIn.Cookie))
	if err != nil {
		return oscar.XMessage{}, err
	}

	return oscar.XMessage{
		SnacFrame: oscar.SnacFrame{
			FoodGroup: oscar.CHAT_NAV,
			SubGroup:  oscar.ChatNavNavInfo,
		},
		SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(0x04, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       4,
						Cookie:         room.Cookie,
						InstanceNumber: 100,
						DetailLevel:    2,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					}),
				},
			},
		},
	}, nil
}
