// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VulndbLeaderElector is an autogenerated mock type for the leaderElector type
type VulndbLeaderElector struct {
	mock.Mock
}

type VulndbLeaderElector_Expecter struct {
	mock *mock.Mock
}

func (_m *VulndbLeaderElector) EXPECT() *VulndbLeaderElector_Expecter {
	return &VulndbLeaderElector_Expecter{mock: &_m.Mock}
}

// IsLeader provides a mock function with given fields:
func (_m *VulndbLeaderElector) IsLeader() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsLeader")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// VulndbLeaderElector_IsLeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsLeader'
type VulndbLeaderElector_IsLeader_Call struct {
	*mock.Call
}

// IsLeader is a helper method to define mock.On call
func (_e *VulndbLeaderElector_Expecter) IsLeader() *VulndbLeaderElector_IsLeader_Call {
	return &VulndbLeaderElector_IsLeader_Call{Call: _e.mock.On("IsLeader")}
}

func (_c *VulndbLeaderElector_IsLeader_Call) Run(run func()) *VulndbLeaderElector_IsLeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VulndbLeaderElector_IsLeader_Call) Return(_a0 bool) *VulndbLeaderElector_IsLeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VulndbLeaderElector_IsLeader_Call) RunAndReturn(run func() bool) *VulndbLeaderElector_IsLeader_Call {
	_c.Call.Return(run)
	return _c
}

// NewVulndbLeaderElector creates a new instance of VulndbLeaderElector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVulndbLeaderElector(t interface {
	mock.TestingT
	Cleanup(func())
}) *VulndbLeaderElector {
	mock := &VulndbLeaderElector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
