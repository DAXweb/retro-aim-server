// Code generated by mockery v2.51.1. DO NOT EDIT.

package toc

import mock "github.com/stretchr/testify/mock"

// mockCookieBaker is an autogenerated mock type for the CookieBaker type
type mockCookieBaker struct {
	mock.Mock
}

type mockCookieBaker_Expecter struct {
	mock *mock.Mock
}

func (_m *mockCookieBaker) EXPECT() *mockCookieBaker_Expecter {
	return &mockCookieBaker_Expecter{mock: &_m.Mock}
}

// Crack provides a mock function with given fields: data
func (_m *mockCookieBaker) Crack(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Crack")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockCookieBaker_Crack_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Crack'
type mockCookieBaker_Crack_Call struct {
	*mock.Call
}

// Crack is a helper method to define mock.On call
//   - data []byte
func (_e *mockCookieBaker_Expecter) Crack(data interface{}) *mockCookieBaker_Crack_Call {
	return &mockCookieBaker_Crack_Call{Call: _e.mock.On("Crack", data)}
}

func (_c *mockCookieBaker_Crack_Call) Run(run func(data []byte)) *mockCookieBaker_Crack_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *mockCookieBaker_Crack_Call) Return(_a0 []byte, _a1 error) *mockCookieBaker_Crack_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockCookieBaker_Crack_Call) RunAndReturn(run func([]byte) ([]byte, error)) *mockCookieBaker_Crack_Call {
	_c.Call.Return(run)
	return _c
}

// Issue provides a mock function with given fields: data
func (_m *mockCookieBaker) Issue(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Issue")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockCookieBaker_Issue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Issue'
type mockCookieBaker_Issue_Call struct {
	*mock.Call
}

// Issue is a helper method to define mock.On call
//   - data []byte
func (_e *mockCookieBaker_Expecter) Issue(data interface{}) *mockCookieBaker_Issue_Call {
	return &mockCookieBaker_Issue_Call{Call: _e.mock.On("Issue", data)}
}

func (_c *mockCookieBaker_Issue_Call) Run(run func(data []byte)) *mockCookieBaker_Issue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *mockCookieBaker_Issue_Call) Return(_a0 []byte, _a1 error) *mockCookieBaker_Issue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockCookieBaker_Issue_Call) RunAndReturn(run func([]byte) ([]byte, error)) *mockCookieBaker_Issue_Call {
	_c.Call.Return(run)
	return _c
}

// newMockCookieBaker creates a new instance of mockCookieBaker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockCookieBaker(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockCookieBaker {
	mock := &mockCookieBaker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
