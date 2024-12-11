// Code generated by mockery v2.51.1. DO NOT EDIT.

package handler

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockICBMService is an autogenerated mock type for the ICBMService type
type mockICBMService struct {
	mock.Mock
}

type mockICBMService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockICBMService) EXPECT() *mockICBMService_Expecter {
	return &mockICBMService_Expecter{mock: &_m.Mock}
}

// ChannelMsgToHost provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockICBMService) ChannelMsgToHost(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*wire.SNACMessage, error) {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for ChannelMsgToHost")
	}

	var r0 *wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*wire.SNACMessage, error)); ok {
		return rf(ctx, sess, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) *wire.SNACMessage); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.SNACMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) error); ok {
		r1 = rf(ctx, sess, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockICBMService_ChannelMsgToHost_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelMsgToHost'
type mockICBMService_ChannelMsgToHost_Call struct {
	*mock.Call
}

// ChannelMsgToHost is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame wire.SNACFrame
//   - inBody wire.SNAC_0x04_0x06_ICBMChannelMsgToHost
func (_e *mockICBMService_Expecter) ChannelMsgToHost(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockICBMService_ChannelMsgToHost_Call {
	return &mockICBMService_ChannelMsgToHost_Call{Call: _e.mock.On("ChannelMsgToHost", ctx, sess, inFrame, inBody)}
}

func (_c *mockICBMService_ChannelMsgToHost_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x06_ICBMChannelMsgToHost)) *mockICBMService_ChannelMsgToHost_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x04_0x06_ICBMChannelMsgToHost))
	})
	return _c
}

func (_c *mockICBMService_ChannelMsgToHost_Call) Return(_a0 *wire.SNACMessage, _a1 error) *mockICBMService_ChannelMsgToHost_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockICBMService_ChannelMsgToHost_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*wire.SNACMessage, error)) *mockICBMService_ChannelMsgToHost_Call {
	_c.Call.Return(run)
	return _c
}

// ClientErr provides a mock function with given fields: ctx, sess, frame, body
func (_m *mockICBMService) ClientErr(ctx context.Context, sess *state.Session, frame wire.SNACFrame, body wire.SNAC_0x04_0x0B_ICBMClientErr) error {
	ret := _m.Called(ctx, sess, frame, body)

	if len(ret) == 0 {
		panic("no return value specified for ClientErr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x0B_ICBMClientErr) error); ok {
		r0 = rf(ctx, sess, frame, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockICBMService_ClientErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientErr'
type mockICBMService_ClientErr_Call struct {
	*mock.Call
}

// ClientErr is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - frame wire.SNACFrame
//   - body wire.SNAC_0x04_0x0B_ICBMClientErr
func (_e *mockICBMService_Expecter) ClientErr(ctx interface{}, sess interface{}, frame interface{}, body interface{}) *mockICBMService_ClientErr_Call {
	return &mockICBMService_ClientErr_Call{Call: _e.mock.On("ClientErr", ctx, sess, frame, body)}
}

func (_c *mockICBMService_ClientErr_Call) Run(run func(ctx context.Context, sess *state.Session, frame wire.SNACFrame, body wire.SNAC_0x04_0x0B_ICBMClientErr)) *mockICBMService_ClientErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x04_0x0B_ICBMClientErr))
	})
	return _c
}

func (_c *mockICBMService_ClientErr_Call) Return(_a0 error) *mockICBMService_ClientErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockICBMService_ClientErr_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x0B_ICBMClientErr) error) *mockICBMService_ClientErr_Call {
	_c.Call.Return(run)
	return _c
}

// ClientEvent provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockICBMService) ClientEvent(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x14_ICBMClientEvent) error {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for ClientEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x14_ICBMClientEvent) error); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockICBMService_ClientEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientEvent'
type mockICBMService_ClientEvent_Call struct {
	*mock.Call
}

// ClientEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame wire.SNACFrame
//   - inBody wire.SNAC_0x04_0x14_ICBMClientEvent
func (_e *mockICBMService_Expecter) ClientEvent(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockICBMService_ClientEvent_Call {
	return &mockICBMService_ClientEvent_Call{Call: _e.mock.On("ClientEvent", ctx, sess, inFrame, inBody)}
}

func (_c *mockICBMService_ClientEvent_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x14_ICBMClientEvent)) *mockICBMService_ClientEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x04_0x14_ICBMClientEvent))
	})
	return _c
}

func (_c *mockICBMService_ClientEvent_Call) Return(_a0 error) *mockICBMService_ClientEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockICBMService_ClientEvent_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x14_ICBMClientEvent) error) *mockICBMService_ClientEvent_Call {
	_c.Call.Return(run)
	return _c
}

// EvilRequest provides a mock function with given fields: ctx, sess, inFrame, inBody
func (_m *mockICBMService) EvilRequest(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x08_ICBMEvilRequest) (wire.SNACMessage, error) {
	ret := _m.Called(ctx, sess, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for EvilRequest")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x08_ICBMEvilRequest) (wire.SNACMessage, error)); ok {
		return rf(ctx, sess, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x08_ICBMEvilRequest) wire.SNACMessage); ok {
		r0 = rf(ctx, sess, inFrame, inBody)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x08_ICBMEvilRequest) error); ok {
		r1 = rf(ctx, sess, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockICBMService_EvilRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EvilRequest'
type mockICBMService_EvilRequest_Call struct {
	*mock.Call
}

// EvilRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame wire.SNACFrame
//   - inBody wire.SNAC_0x04_0x08_ICBMEvilRequest
func (_e *mockICBMService_Expecter) EvilRequest(ctx interface{}, sess interface{}, inFrame interface{}, inBody interface{}) *mockICBMService_EvilRequest_Call {
	return &mockICBMService_EvilRequest_Call{Call: _e.mock.On("EvilRequest", ctx, sess, inFrame, inBody)}
}

func (_c *mockICBMService_EvilRequest_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x08_ICBMEvilRequest)) *mockICBMService_EvilRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x04_0x08_ICBMEvilRequest))
	})
	return _c
}

func (_c *mockICBMService_EvilRequest_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockICBMService_EvilRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockICBMService_EvilRequest_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x04_0x08_ICBMEvilRequest) (wire.SNACMessage, error)) *mockICBMService_EvilRequest_Call {
	_c.Call.Return(run)
	return _c
}

// ParameterQuery provides a mock function with given fields: ctx, inFrame
func (_m *mockICBMService) ParameterQuery(ctx context.Context, inFrame wire.SNACFrame) wire.SNACMessage {
	ret := _m.Called(ctx, inFrame)

	if len(ret) == 0 {
		panic("no return value specified for ParameterQuery")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(ctx, inFrame)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockICBMService_ParameterQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParameterQuery'
type mockICBMService_ParameterQuery_Call struct {
	*mock.Call
}

// ParameterQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - inFrame wire.SNACFrame
func (_e *mockICBMService_Expecter) ParameterQuery(ctx interface{}, inFrame interface{}) *mockICBMService_ParameterQuery_Call {
	return &mockICBMService_ParameterQuery_Call{Call: _e.mock.On("ParameterQuery", ctx, inFrame)}
}

func (_c *mockICBMService_ParameterQuery_Call) Run(run func(ctx context.Context, inFrame wire.SNACFrame)) *mockICBMService_ParameterQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockICBMService_ParameterQuery_Call) Return(_a0 wire.SNACMessage) *mockICBMService_ParameterQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockICBMService_ParameterQuery_Call) RunAndReturn(run func(context.Context, wire.SNACFrame) wire.SNACMessage) *mockICBMService_ParameterQuery_Call {
	_c.Call.Return(run)
	return _c
}

// newMockICBMService creates a new instance of mockICBMService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockICBMService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockICBMService {
	mock := &mockICBMService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
