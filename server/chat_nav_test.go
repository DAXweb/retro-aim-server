package server

import (
	"bytes"
	"context"
	"github.com/mkaminski/goaim/user"
	"log/slog"
	"testing"
	"time"

	"github.com/mkaminski/goaim/oscar"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSendAndReceiveCreateRoom(t *testing.T) {
	//
	// build dependencies
	//
	userSess := newTestSession("user-screen-name", sessOptCannedID)

	cr := NewChatRegistry()

	sm := NewMockSessionManager(t)
	sm.EXPECT().NewSessionWithSN(userSess.ID(), userSess.ScreenName()).
		Return(&user.Session{})

	crf := func(logger *slog.Logger) ChatRoom {
		return ChatRoom{
			Cookie:         "dummy-cookie",
			CreateTime:     time.UnixMilli(0),
			SessionManager: sm,
		}
	}

	//
	// send input SNAC
	//
	inputSNAC := oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
		Exchange:       1,
		Cookie:         "create", // actual canned value sent by AIM client
		InstanceNumber: 2,
		DetailLevel:    3,
		TLVBlock: oscar.TLVBlock{
			TLVList: oscar.TLVList{
				oscar.NewTLV(oscar.ChatTLVRoomName, "the-chat-room-name"),
			},
		},
	}
	svc := ChatNavService{
		cr: cr,
	}
	outputSNAC, err := svc.CreateRoomHandler(context.Background(), userSess, crf, inputSNAC)
	assert.NoError(t, err)

	//
	// verify chat room created by handler
	//
	expectChatRoom := ChatRoom{
		SessionManager: sm,
		Cookie:         "dummy-cookie",
		CreateTime:     time.UnixMilli(0),
		DetailLevel:    3,
		Exchange:       1,
		InstanceNumber: 2,
		Name:           "the-chat-room-name",
	}
	chatRoom, err := cr.Retrieve("dummy-cookie")
	assert.NoError(t, err)
	assert.Equal(t, expectChatRoom, chatRoom)

	//
	// send input SNAC
	//
	expectSNAC := oscar.XMessage{
		SnacFrame: oscar.SnacFrame{
			FoodGroup: oscar.CHAT_NAV,
			SubGroup:  oscar.ChatNavNavInfo,
		},
		SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(
						oscar.ChatNavTLVRoomInfo,
						oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
							Exchange:       chatRoom.Exchange,
							Cookie:         chatRoom.Cookie,
							InstanceNumber: chatRoom.InstanceNumber,
							DetailLevel:    chatRoom.DetailLevel,
							TLVBlock: oscar.TLVBlock{
								TLVList: chatRoom.TLVList(),
							},
						},
					),
				},
			},
		},
	}

	assert.Equal(t, expectSNAC, outputSNAC)
}

func TestChatNavRouter_RouteChatNavRouter(t *testing.T) {
	cases := []struct {
		// name is the unit test name
		name string
		// input is the request payload
		input oscar.XMessage
		// output is the response payload
		output oscar.XMessage
		// handlerErr is the mocked handler error response
		handlerErr error
		// expectErr is the expected error returned by the router
		expectErr error
	}{
		{
			name: "receive ChatNavRequestChatRights, return ChatNavNavInfo",
			input: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavRequestChatRights,
				},
				SnacOut: struct{}{},
			},
			output: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavNavInfo,
				},
				SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
					TLVRestBlock: oscar.TLVRestBlock{
						TLVList: oscar.TLVList{
							oscar.NewTLV(0x02, uint8(10)),
						},
					},
				},
			},
		},
		{
			name: "receive ChatNavRequestRoomInfo, return ChatNavNavInfo",
			input: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavRequestRoomInfo,
				},
				SnacOut: oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo{
					Exchange: 1,
				},
			},
			output: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavNavInfo,
				},
				SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
					TLVRestBlock: oscar.TLVRestBlock{
						TLVList: oscar.TLVList{
							oscar.NewTLV(0x02, uint8(10)),
						},
					},
				},
			},
		},
		{
			name: "receive ChatNavCreateRoom, return ChatNavNavInfo",
			input: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavCreateRoom,
				},
				SnacOut: oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
					Exchange: 1,
				},
			},
			output: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavNavInfo,
				},
				SnacOut: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
					TLVRestBlock: oscar.TLVRestBlock{
						TLVList: oscar.TLVList{
							oscar.NewTLV(0x02, uint8(10)),
						},
					},
				},
			},
		},
		{
			name: "receive ChatNavRequestOccupantList, return ErrUnsupportedSubGroup",
			input: oscar.XMessage{
				SnacFrame: oscar.SnacFrame{
					FoodGroup: oscar.CHAT_NAV,
					SubGroup:  oscar.ChatNavRequestOccupantList,
				},
				SnacOut: struct{}{},
			},
			output:    oscar.XMessage{},
			expectErr: ErrUnsupportedSubGroup,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewMockChatNavHandler(t)
			svc.EXPECT().
				RequestChatRightsHandler(mock.Anything).
				Return(tc.output).
				Maybe()
			svc.EXPECT().
				RequestRoomInfoHandler(mock.Anything, tc.input.SnacOut).
				Return(tc.output, tc.handlerErr).
				Maybe()
			svc.EXPECT().
				CreateRoomHandler(mock.Anything, mock.Anything, mock.Anything, tc.input.SnacOut).
				Return(tc.output, tc.handlerErr).
				Maybe()

			router := ChatNavRouter{
				ChatNavHandler: svc,
				RouteLogger: RouteLogger{
					Logger: NewLogger(Config{}),
				},
			}

			bufIn := &bytes.Buffer{}
			assert.NoError(t, oscar.Marshal(tc.input.SnacOut, bufIn))

			bufOut := &bytes.Buffer{}
			seq := uint32(0)

			err := router.RouteChatNav(nil, nil, tc.input.SnacFrame, bufIn, bufOut, &seq)
			assert.ErrorIs(t, err, tc.expectErr)
			if tc.expectErr != nil {
				return
			}

			if tc.output.SnacFrame == (oscar.SnacFrame{}) {
				return
			}

			// verify the FLAP frame
			flap := oscar.FlapFrame{}
			assert.NoError(t, oscar.Unmarshal(&flap, bufOut))

			// make sure the sequence increments
			assert.Equal(t, seq, uint32(1))
			assert.Equal(t, flap.Sequence, uint16(0))

			flapBuf, err := flap.SNACBuffer(bufOut)
			assert.NoError(t, err)

			// verify the SNAC frame
			snacFrame := oscar.SnacFrame{}
			assert.NoError(t, oscar.Unmarshal(&snacFrame, flapBuf))
			assert.Equal(t, tc.output.SnacFrame, snacFrame)

			// verify the SNAC message
			snacBuf := &bytes.Buffer{}
			assert.NoError(t, oscar.Marshal(tc.output.SnacOut, snacBuf))
			assert.Equal(t, snacBuf.Bytes(), flapBuf.Bytes())
		})
	}
}
