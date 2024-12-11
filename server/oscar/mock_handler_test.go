// Code generated by mockery v2.51.1. DO NOT EDIT.

package oscar

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	state "github.com/mk6i/retro-aim-server/state"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockHandler is an autogenerated mock type for the Handler type
type mockHandler struct {
	mock.Mock
}

type mockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockHandler) EXPECT() *mockHandler_Expecter {
	return &mockHandler_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, sess, inFrame, r, rw
func (_m *mockHandler) Handle(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, r io.Reader, rw ResponseWriter) error {
	ret := _m.Called(ctx, sess, inFrame, r, rw)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, wire.SNACFrame, io.Reader, ResponseWriter) error); ok {
		r0 = rf(ctx, sess, inFrame, r, rw)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type mockHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - inFrame wire.SNACFrame
//   - r io.Reader
//   - rw ResponseWriter
func (_e *mockHandler_Expecter) Handle(ctx interface{}, sess interface{}, inFrame interface{}, r interface{}, rw interface{}) *mockHandler_Handle_Call {
	return &mockHandler_Handle_Call{Call: _e.mock.On("Handle", ctx, sess, inFrame, r, rw)}
}

func (_c *mockHandler_Handle_Call) Run(run func(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, r io.Reader, rw ResponseWriter)) *mockHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(wire.SNACFrame), args[3].(io.Reader), args[4].(ResponseWriter))
	})
	return _c
}

func (_c *mockHandler_Handle_Call) Return(_a0 error) *mockHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockHandler_Handle_Call) RunAndReturn(run func(context.Context, *state.Session, wire.SNACFrame, io.Reader, ResponseWriter) error) *mockHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// newMockHandler creates a new instance of mockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockHandler {
	mock := &mockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
