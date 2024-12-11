// Code generated by mockery v2.51.1. DO NOT EDIT.

package foodgroup

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockLocalBuddyListManager is an autogenerated mock type for the LocalBuddyListManager type
type mockLocalBuddyListManager struct {
	mock.Mock
}

type mockLocalBuddyListManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockLocalBuddyListManager) EXPECT() *mockLocalBuddyListManager_Expecter {
	return &mockLocalBuddyListManager_Expecter{mock: &_m.Mock}
}

// AddBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) AddBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for AddBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_AddBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBuddy'
type mockLocalBuddyListManager_AddBuddy_Call struct {
	*mock.Call
}

// AddBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) AddBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_AddBuddy_Call {
	return &mockLocalBuddyListManager_AddBuddy_Call{Call: _e.mock.On("AddBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_AddBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_AddBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_AddBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_AddBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_AddBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_AddBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// DenyBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) DenyBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for DenyBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_DenyBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DenyBuddy'
type mockLocalBuddyListManager_DenyBuddy_Call struct {
	*mock.Call
}

// DenyBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) DenyBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_DenyBuddy_Call {
	return &mockLocalBuddyListManager_DenyBuddy_Call{Call: _e.mock.On("DenyBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_DenyBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_DenyBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_DenyBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_DenyBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_DenyBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_DenyBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// PermitBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) PermitBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for PermitBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_PermitBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PermitBuddy'
type mockLocalBuddyListManager_PermitBuddy_Call struct {
	*mock.Call
}

// PermitBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) PermitBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_PermitBuddy_Call {
	return &mockLocalBuddyListManager_PermitBuddy_Call{Call: _e.mock.On("PermitBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_PermitBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_PermitBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_PermitBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_PermitBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_PermitBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_PermitBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) RemoveBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for RemoveBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_RemoveBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveBuddy'
type mockLocalBuddyListManager_RemoveBuddy_Call struct {
	*mock.Call
}

// RemoveBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) RemoveBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_RemoveBuddy_Call {
	return &mockLocalBuddyListManager_RemoveBuddy_Call{Call: _e.mock.On("RemoveBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_RemoveBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_RemoveBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_RemoveBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_RemoveBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_RemoveBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_RemoveBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveDenyBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) RemoveDenyBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for RemoveDenyBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_RemoveDenyBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveDenyBuddy'
type mockLocalBuddyListManager_RemoveDenyBuddy_Call struct {
	*mock.Call
}

// RemoveDenyBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) RemoveDenyBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_RemoveDenyBuddy_Call {
	return &mockLocalBuddyListManager_RemoveDenyBuddy_Call{Call: _e.mock.On("RemoveDenyBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_RemoveDenyBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_RemoveDenyBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_RemoveDenyBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_RemoveDenyBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_RemoveDenyBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_RemoveDenyBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// RemovePermitBuddy provides a mock function with given fields: me, them
func (_m *mockLocalBuddyListManager) RemovePermitBuddy(me state.IdentScreenName, them state.IdentScreenName) error {
	ret := _m.Called(me, them)

	if len(ret) == 0 {
		panic("no return value specified for RemovePermitBuddy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, state.IdentScreenName) error); ok {
		r0 = rf(me, them)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_RemovePermitBuddy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemovePermitBuddy'
type mockLocalBuddyListManager_RemovePermitBuddy_Call struct {
	*mock.Call
}

// RemovePermitBuddy is a helper method to define mock.On call
//   - me state.IdentScreenName
//   - them state.IdentScreenName
func (_e *mockLocalBuddyListManager_Expecter) RemovePermitBuddy(me interface{}, them interface{}) *mockLocalBuddyListManager_RemovePermitBuddy_Call {
	return &mockLocalBuddyListManager_RemovePermitBuddy_Call{Call: _e.mock.On("RemovePermitBuddy", me, them)}
}

func (_c *mockLocalBuddyListManager_RemovePermitBuddy_Call) Run(run func(me state.IdentScreenName, them state.IdentScreenName)) *mockLocalBuddyListManager_RemovePermitBuddy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_RemovePermitBuddy_Call) Return(_a0 error) *mockLocalBuddyListManager_RemovePermitBuddy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_RemovePermitBuddy_Call) RunAndReturn(run func(state.IdentScreenName, state.IdentScreenName) error) *mockLocalBuddyListManager_RemovePermitBuddy_Call {
	_c.Call.Return(run)
	return _c
}

// SetPDMode provides a mock function with given fields: user, pdMode
func (_m *mockLocalBuddyListManager) SetPDMode(user state.IdentScreenName, pdMode wire.FeedbagPDMode) error {
	ret := _m.Called(user, pdMode)

	if len(ret) == 0 {
		panic("no return value specified for SetPDMode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName, wire.FeedbagPDMode) error); ok {
		r0 = rf(user, pdMode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockLocalBuddyListManager_SetPDMode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPDMode'
type mockLocalBuddyListManager_SetPDMode_Call struct {
	*mock.Call
}

// SetPDMode is a helper method to define mock.On call
//   - user state.IdentScreenName
//   - pdMode wire.FeedbagPDMode
func (_e *mockLocalBuddyListManager_Expecter) SetPDMode(user interface{}, pdMode interface{}) *mockLocalBuddyListManager_SetPDMode_Call {
	return &mockLocalBuddyListManager_SetPDMode_Call{Call: _e.mock.On("SetPDMode", user, pdMode)}
}

func (_c *mockLocalBuddyListManager_SetPDMode_Call) Run(run func(user state.IdentScreenName, pdMode wire.FeedbagPDMode)) *mockLocalBuddyListManager_SetPDMode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName), args[1].(wire.FeedbagPDMode))
	})
	return _c
}

func (_c *mockLocalBuddyListManager_SetPDMode_Call) Return(_a0 error) *mockLocalBuddyListManager_SetPDMode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockLocalBuddyListManager_SetPDMode_Call) RunAndReturn(run func(state.IdentScreenName, wire.FeedbagPDMode) error) *mockLocalBuddyListManager_SetPDMode_Call {
	_c.Call.Return(run)
	return _c
}

// newMockLocalBuddyListManager creates a new instance of mockLocalBuddyListManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockLocalBuddyListManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockLocalBuddyListManager {
	mock := &mockLocalBuddyListManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
