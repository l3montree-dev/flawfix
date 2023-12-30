// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// CoreProject is an autogenerated mock type for the Project type
type CoreProject struct {
	mock.Mock
}

type CoreProject_Expecter struct {
	mock *mock.Mock
}

func (_m *CoreProject) EXPECT() *CoreProject_Expecter {
	return &CoreProject_Expecter{mock: &_m.Mock}
}

// GetID provides a mock function with given fields:
func (_m *CoreProject) GetID() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// CoreProject_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type CoreProject_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
func (_e *CoreProject_Expecter) GetID() *CoreProject_GetID_Call {
	return &CoreProject_GetID_Call{Call: _e.mock.On("GetID")}
}

func (_c *CoreProject_GetID_Call) Run(run func()) *CoreProject_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CoreProject_GetID_Call) Return(_a0 uuid.UUID) *CoreProject_GetID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CoreProject_GetID_Call) RunAndReturn(run func() uuid.UUID) *CoreProject_GetID_Call {
	_c.Call.Return(run)
	return _c
}

// NewCoreProject creates a new instance of CoreProject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCoreProject(t interface {
	mock.TestingT
	Cleanup(func())
}) *CoreProject {
	mock := &CoreProject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}