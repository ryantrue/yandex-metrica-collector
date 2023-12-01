// Code generated by mockery. DO NOT EDIT.

package server

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockServer is an autogenerated mock type for the server type
type mockServer struct {
	mock.Mock
}

// ListenAndServe provides a mock function with given fields:
func (_m *mockServer) ListenAndServe() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with given fields: ctx
func (_m *mockServer) Shutdown(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockServer interface {
	mock.TestingT
	Cleanup(func())
}

// newMockServer creates a new instance of mockServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockServer(t mockConstructorTestingTnewMockServer) *mockServer {
	mock := &mockServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}