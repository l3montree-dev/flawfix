// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/devguard/internal/database/models"
	mock "github.com/stretchr/testify/mock"
)

// AssetVulnService is an autogenerated mock type for the vulnService type
type AssetVulnService struct {
	mock.Mock
}

type AssetVulnService_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetVulnService) EXPECT() *AssetVulnService_Expecter {
	return &AssetVulnService_Expecter{mock: &_m.Mock}
}

// GetVulnsForAll provides a mock function with given fields: purls
func (_m *AssetVulnService) GetVulnsForAll(purls []string) ([]models.VulnInPackage, error) {
	ret := _m.Called(purls)

	if len(ret) == 0 {
		panic("no return value specified for GetVulnsForAll")
	}

	var r0 []models.VulnInPackage
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]models.VulnInPackage, error)); ok {
		return rf(purls)
	}
	if rf, ok := ret.Get(0).(func([]string) []models.VulnInPackage); ok {
		r0 = rf(purls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.VulnInPackage)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(purls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetVulnService_GetVulnsForAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVulnsForAll'
type AssetVulnService_GetVulnsForAll_Call struct {
	*mock.Call
}

// GetVulnsForAll is a helper method to define mock.On call
//   - purls []string
func (_e *AssetVulnService_Expecter) GetVulnsForAll(purls interface{}) *AssetVulnService_GetVulnsForAll_Call {
	return &AssetVulnService_GetVulnsForAll_Call{Call: _e.mock.On("GetVulnsForAll", purls)}
}

func (_c *AssetVulnService_GetVulnsForAll_Call) Run(run func(purls []string)) *AssetVulnService_GetVulnsForAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *AssetVulnService_GetVulnsForAll_Call) Return(_a0 []models.VulnInPackage, _a1 error) *AssetVulnService_GetVulnsForAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetVulnService_GetVulnsForAll_Call) RunAndReturn(run func([]string) ([]models.VulnInPackage, error)) *AssetVulnService_GetVulnsForAll_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssetVulnService creates a new instance of AssetVulnService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssetVulnService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AssetVulnService {
	mock := &AssetVulnService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
