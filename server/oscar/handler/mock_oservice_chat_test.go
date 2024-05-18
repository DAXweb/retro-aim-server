// Code generated by mockery v2.40.1. DO NOT EDIT.

package handler

import (
	context "context"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockOServiceChatService is an autogenerated mock type for the OServiceChatService type
type mockOServiceChatService struct {
	mock.Mock
}

type mockOServiceChatService_Expecter struct {
	mock *mock.Mock
}

func (_m *mockOServiceChatService) EXPECT() *mockOServiceChatService_Expecter {
	return &mockOServiceChatService_Expecter{mock: &_m.Mock}
}

// ClientOnline provides a mock function with given fields: ctx, sess
func (_m *mockOServiceChatService) ClientOnline(ctx context.Context, sess *state.Session) error {
	ret := _m.Called(ctx, sess)

	if len(ret) == 0 {
		panic("no return value specified for ClientOnline")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session) error); ok {
		r0 = rf(ctx, sess)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOServiceChatService_ClientOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientOnline'
type mockOServiceChatService_ClientOnline_Call struct {
	*mock.Call
}

// ClientOnline is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
func (_e *mockOServiceChatService_Expecter) ClientOnline(ctx interface{}, sess interface{}) *mockOServiceChatService_ClientOnline_Call {
	return &mockOServiceChatService_ClientOnline_Call{Call: _e.mock.On("ClientOnline", ctx, sess)}
}

func (_c *mockOServiceChatService_ClientOnline_Call) Run(run func(ctx context.Context, sess *state.Session)) *mockOServiceChatService_ClientOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session))
	})
	return _c
}

func (_c *mockOServiceChatService_ClientOnline_Call) Return(_a0 error) *mockOServiceChatService_ClientOnline_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_ClientOnline_Call) RunAndReturn(run func(context.Context, *state.Session) error) *mockOServiceChatService_ClientOnline_Call {
	_c.Call.Return(run)
	return _c
}

// ClientVersions provides a mock function with given fields: ctx, frame, bodyIn
func (_m *mockOServiceChatService) ClientVersions(ctx context.Context, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage {
	ret := _m.Called(ctx, frame, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for ClientVersions")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame, wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage); ok {
		r0 = rf(ctx, frame, bodyIn)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceChatService_ClientVersions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClientVersions'
type mockOServiceChatService_ClientVersions_Call struct {
	*mock.Call
}

// ClientVersions is a helper method to define mock.On call
//   - ctx context.Context
//   - frame wire.SNACFrame
//   - bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions
func (_e *mockOServiceChatService_Expecter) ClientVersions(ctx interface{}, frame interface{}, bodyIn interface{}) *mockOServiceChatService_ClientVersions_Call {
	return &mockOServiceChatService_ClientVersions_Call{Call: _e.mock.On("ClientVersions", ctx, frame, bodyIn)}
}

func (_c *mockOServiceChatService_ClientVersions_Call) Run(run func(ctx context.Context, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x17_OServiceClientVersions)) *mockOServiceChatService_ClientVersions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame), args[2].(wire.SNAC_0x01_0x17_OServiceClientVersions))
	})
	return _c
}

func (_c *mockOServiceChatService_ClientVersions_Call) Return(_a0 wire.SNACMessage) *mockOServiceChatService_ClientVersions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_ClientVersions_Call) RunAndReturn(run func(context.Context, wire.SNACFrame, wire.SNAC_0x01_0x17_OServiceClientVersions) wire.SNACMessage) *mockOServiceChatService_ClientVersions_Call {
	_c.Call.Return(run)
	return _c
}

// HostOnline provides a mock function with given fields:
func (_m *mockOServiceChatService) HostOnline() wire.SNACMessage {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HostOnline")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func() wire.SNACMessage); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceChatService_HostOnline_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HostOnline'
type mockOServiceChatService_HostOnline_Call struct {
	*mock.Call
}

// HostOnline is a helper method to define mock.On call
func (_e *mockOServiceChatService_Expecter) HostOnline() *mockOServiceChatService_HostOnline_Call {
	return &mockOServiceChatService_HostOnline_Call{Call: _e.mock.On("HostOnline")}
}

func (_c *mockOServiceChatService_HostOnline_Call) Run(run func()) *mockOServiceChatService_HostOnline_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockOServiceChatService_HostOnline_Call) Return(_a0 wire.SNACMessage) *mockOServiceChatService_HostOnline_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_HostOnline_Call) RunAndReturn(run func() wire.SNACMessage) *mockOServiceChatService_HostOnline_Call {
	_c.Call.Return(run)
	return _c
}

// IdleNotification provides a mock function with given fields: ctx, sess, bodyIn
func (_m *mockOServiceChatService) IdleNotification(ctx context.Context, sess *state.Session, bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification) error {
	ret := _m.Called(ctx, sess, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for IdleNotification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNAC_0x01_0x11_OServiceIdleNotification) error); ok {
		r0 = rf(ctx, sess, bodyIn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockOServiceChatService_IdleNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IdleNotification'
type mockOServiceChatService_IdleNotification_Call struct {
	*mock.Call
}

// IdleNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification
func (_e *mockOServiceChatService_Expecter) IdleNotification(ctx interface{}, sess interface{}, bodyIn interface{}) *mockOServiceChatService_IdleNotification_Call {
	return &mockOServiceChatService_IdleNotification_Call{Call: _e.mock.On("IdleNotification", ctx, sess, bodyIn)}
}

func (_c *mockOServiceChatService_IdleNotification_Call) Run(run func(ctx context.Context, sess *state.Session, bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification)) *mockOServiceChatService_IdleNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNAC_0x01_0x11_OServiceIdleNotification))
	})
	return _c
}

func (_c *mockOServiceChatService_IdleNotification_Call) Return(_a0 error) *mockOServiceChatService_IdleNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_IdleNotification_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNAC_0x01_0x11_OServiceIdleNotification) error) *mockOServiceChatService_IdleNotification_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsQuery provides a mock function with given fields: ctx, frame
func (_m *mockOServiceChatService) RateParamsQuery(ctx context.Context, frame wire.SNACFrame) wire.SNACMessage {
	ret := _m.Called(ctx, frame)

	if len(ret) == 0 {
		panic("no return value specified for RateParamsQuery")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(ctx, frame)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceChatService_RateParamsQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsQuery'
type mockOServiceChatService_RateParamsQuery_Call struct {
	*mock.Call
}

// RateParamsQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - frame wire.SNACFrame
func (_e *mockOServiceChatService_Expecter) RateParamsQuery(ctx interface{}, frame interface{}) *mockOServiceChatService_RateParamsQuery_Call {
	return &mockOServiceChatService_RateParamsQuery_Call{Call: _e.mock.On("RateParamsQuery", ctx, frame)}
}

func (_c *mockOServiceChatService_RateParamsQuery_Call) Run(run func(ctx context.Context, frame wire.SNACFrame)) *mockOServiceChatService_RateParamsQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockOServiceChatService_RateParamsQuery_Call) Return(_a0 wire.SNACMessage) *mockOServiceChatService_RateParamsQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_RateParamsQuery_Call) RunAndReturn(run func(context.Context, wire.SNACFrame) wire.SNACMessage) *mockOServiceChatService_RateParamsQuery_Call {
	_c.Call.Return(run)
	return _c
}

// RateParamsSubAdd provides a mock function with given fields: _a0, _a1
func (_m *mockOServiceChatService) RateParamsSubAdd(_a0 context.Context, _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd) {
	_m.Called(_a0, _a1)
}

// mockOServiceChatService_RateParamsSubAdd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RateParamsSubAdd'
type mockOServiceChatService_RateParamsSubAdd_Call struct {
	*mock.Call
}

// RateParamsSubAdd is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd
func (_e *mockOServiceChatService_Expecter) RateParamsSubAdd(_a0 interface{}, _a1 interface{}) *mockOServiceChatService_RateParamsSubAdd_Call {
	return &mockOServiceChatService_RateParamsSubAdd_Call{Call: _e.mock.On("RateParamsSubAdd", _a0, _a1)}
}

func (_c *mockOServiceChatService_RateParamsSubAdd_Call) Run(run func(_a0 context.Context, _a1 wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceChatService_RateParamsSubAdd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd))
	})
	return _c
}

func (_c *mockOServiceChatService_RateParamsSubAdd_Call) Return() *mockOServiceChatService_RateParamsSubAdd_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockOServiceChatService_RateParamsSubAdd_Call) RunAndReturn(run func(context.Context, wire.SNAC_0x01_0x08_OServiceRateParamsSubAdd)) *mockOServiceChatService_RateParamsSubAdd_Call {
	_c.Call.Return(run)
	return _c
}

// SetPrivacyFlags provides a mock function with given fields: ctx, bodyIn
func (_m *mockOServiceChatService) SetPrivacyFlags(ctx context.Context, bodyIn wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags) {
	_m.Called(ctx, bodyIn)
}

// mockOServiceChatService_SetPrivacyFlags_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPrivacyFlags'
type mockOServiceChatService_SetPrivacyFlags_Call struct {
	*mock.Call
}

// SetPrivacyFlags is a helper method to define mock.On call
//   - ctx context.Context
//   - bodyIn wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags
func (_e *mockOServiceChatService_Expecter) SetPrivacyFlags(ctx interface{}, bodyIn interface{}) *mockOServiceChatService_SetPrivacyFlags_Call {
	return &mockOServiceChatService_SetPrivacyFlags_Call{Call: _e.mock.On("SetPrivacyFlags", ctx, bodyIn)}
}

func (_c *mockOServiceChatService_SetPrivacyFlags_Call) Run(run func(ctx context.Context, bodyIn wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags)) *mockOServiceChatService_SetPrivacyFlags_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags))
	})
	return _c
}

func (_c *mockOServiceChatService_SetPrivacyFlags_Call) Return() *mockOServiceChatService_SetPrivacyFlags_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockOServiceChatService_SetPrivacyFlags_Call) RunAndReturn(run func(context.Context, wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags)) *mockOServiceChatService_SetPrivacyFlags_Call {
	_c.Call.Return(run)
	return _c
}

// SetUserInfoFields provides a mock function with given fields: ctx, sess, frame, bodyIn
func (_m *mockOServiceChatService) SetUserInfoFields(ctx context.Context, sess *state.Session, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error) {
	ret := _m.Called(ctx, sess, frame, bodyIn)

	if len(ret) == 0 {
		panic("no return value specified for SetUserInfoFields")
	}

	var r0 wire.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error)); ok {
		return rf(ctx, sess, frame, bodyIn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) wire.SNACMessage); ok {
		r0 = rf(ctx, sess, frame, bodyIn)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) error); ok {
		r1 = rf(ctx, sess, frame, bodyIn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockOServiceChatService_SetUserInfoFields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetUserInfoFields'
type mockOServiceChatService_SetUserInfoFields_Call struct {
	*mock.Call
}

// SetUserInfoFields is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - frame wire.SNACFrame
//   - bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields
func (_e *mockOServiceChatService_Expecter) SetUserInfoFields(ctx interface{}, sess interface{}, frame interface{}, bodyIn interface{}) *mockOServiceChatService_SetUserInfoFields_Call {
	return &mockOServiceChatService_SetUserInfoFields_Call{Call: _e.mock.On("SetUserInfoFields", ctx, sess, frame, bodyIn)}
}

func (_c *mockOServiceChatService_SetUserInfoFields_Call) Run(run func(ctx context.Context, sess *state.Session, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields)) *mockOServiceChatService_SetUserInfoFields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
	})
	return _c
}

func (_c *mockOServiceChatService_SetUserInfoFields_Call) Return(_a0 wire.SNACMessage, _a1 error) *mockOServiceChatService_SetUserInfoFields_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockOServiceChatService_SetUserInfoFields_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields) (wire.SNACMessage, error)) *mockOServiceChatService_SetUserInfoFields_Call {
	_c.Call.Return(run)
	return _c
}

// UserInfoQuery provides a mock function with given fields: ctx, sess, frame
func (_m *mockOServiceChatService) UserInfoQuery(ctx context.Context, sess *state.Session, frame wire.SNACFrame) wire.SNACMessage {
	ret := _m.Called(ctx, sess, frame)

	if len(ret) == 0 {
		panic("no return value specified for UserInfoQuery")
	}

	var r0 wire.SNACMessage
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame) wire.SNACMessage); ok {
		r0 = rf(ctx, sess, frame)
	} else {
		r0 = ret.Get(0).(wire.SNACMessage)
	}

	return r0
}

// mockOServiceChatService_UserInfoQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserInfoQuery'
type mockOServiceChatService_UserInfoQuery_Call struct {
	*mock.Call
}

// UserInfoQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - frame wire.SNACFrame
func (_e *mockOServiceChatService_Expecter) UserInfoQuery(ctx interface{}, sess interface{}, frame interface{}) *mockOServiceChatService_UserInfoQuery_Call {
	return &mockOServiceChatService_UserInfoQuery_Call{Call: _e.mock.On("UserInfoQuery", ctx, sess, frame)}
}

func (_c *mockOServiceChatService_UserInfoQuery_Call) Run(run func(ctx context.Context, sess *state.Session, frame wire.SNACFrame)) *mockOServiceChatService_UserInfoQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame))
	})
	return _c
}

func (_c *mockOServiceChatService_UserInfoQuery_Call) Return(_a0 wire.SNACMessage) *mockOServiceChatService_UserInfoQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockOServiceChatService_UserInfoQuery_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame) wire.SNACMessage) *mockOServiceChatService_UserInfoQuery_Call {
	_c.Call.Return(run)
	return _c
}

// newMockOServiceChatService creates a new instance of mockOServiceChatService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockOServiceChatService(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockOServiceChatService {
	mock := &mockOServiceChatService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
