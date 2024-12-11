// Code generated by mockery v2.51.1. DO NOT EDIT.

package http

import (
	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"

	wire "github.com/mk6i/retro-aim-server/wire"
)

// mockFeedBagRetriever is an autogenerated mock type for the FeedBagRetriever type
type mockFeedBagRetriever struct {
	mock.Mock
}

type mockFeedBagRetriever_Expecter struct {
	mock *mock.Mock
}

func (_m *mockFeedBagRetriever) EXPECT() *mockFeedBagRetriever_Expecter {
	return &mockFeedBagRetriever_Expecter{mock: &_m.Mock}
}

// BuddyIconRefByName provides a mock function with given fields: screenName
func (_m *mockFeedBagRetriever) BuddyIconRefByName(screenName state.IdentScreenName) (*wire.BARTID, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for BuddyIconRefByName")
	}

	var r0 *wire.BARTID
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (*wire.BARTID, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) *wire.BARTID); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*wire.BARTID)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFeedBagRetriever_BuddyIconRefByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuddyIconRefByName'
type mockFeedBagRetriever_BuddyIconRefByName_Call struct {
	*mock.Call
}

// BuddyIconRefByName is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockFeedBagRetriever_Expecter) BuddyIconRefByName(screenName interface{}) *mockFeedBagRetriever_BuddyIconRefByName_Call {
	return &mockFeedBagRetriever_BuddyIconRefByName_Call{Call: _e.mock.On("BuddyIconRefByName", screenName)}
}

func (_c *mockFeedBagRetriever_BuddyIconRefByName_Call) Run(run func(screenName state.IdentScreenName)) *mockFeedBagRetriever_BuddyIconRefByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockFeedBagRetriever_BuddyIconRefByName_Call) Return(_a0 *wire.BARTID, _a1 error) *mockFeedBagRetriever_BuddyIconRefByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFeedBagRetriever_BuddyIconRefByName_Call) RunAndReturn(run func(state.IdentScreenName) (*wire.BARTID, error)) *mockFeedBagRetriever_BuddyIconRefByName_Call {
	_c.Call.Return(run)
	return _c
}

// newMockFeedBagRetriever creates a new instance of mockFeedBagRetriever. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockFeedBagRetriever(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockFeedBagRetriever {
	mock := &mockFeedBagRetriever{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
