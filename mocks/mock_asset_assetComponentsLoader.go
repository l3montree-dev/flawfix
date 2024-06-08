// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// AssetAssetComponentsLoader is an autogenerated mock type for the assetComponentsLoader type
type AssetAssetComponentsLoader struct {
	mock.Mock
}

type AssetAssetComponentsLoader_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetAssetComponentsLoader) EXPECT() *AssetAssetComponentsLoader_Expecter {
	return &AssetAssetComponentsLoader_Expecter{mock: &_m.Mock}
}

// LoadAssetComponents provides a mock function with given fields: tx, _a1, version
func (_m *AssetAssetComponentsLoader) LoadAssetComponents(tx *gorm.DB, _a1 models.Asset, version string) ([]models.ComponentDependency, error) {
	ret := _m.Called(tx, _a1, version)

	if len(ret) == 0 {
		panic("no return value specified for LoadAssetComponents")
	}

	var r0 []models.ComponentDependency
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, models.Asset, string) ([]models.ComponentDependency, error)); ok {
		return rf(tx, _a1, version)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, models.Asset, string) []models.ComponentDependency); ok {
		r0 = rf(tx, _a1, version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ComponentDependency)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, models.Asset, string) error); ok {
		r1 = rf(tx, _a1, version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetAssetComponentsLoader_LoadAssetComponents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadAssetComponents'
type AssetAssetComponentsLoader_LoadAssetComponents_Call struct {
	*mock.Call
}

// LoadAssetComponents is a helper method to define mock.On call
//   - tx *gorm.DB
//   - _a1 models.Asset
//   - version string
func (_e *AssetAssetComponentsLoader_Expecter) LoadAssetComponents(tx interface{}, _a1 interface{}, version interface{}) *AssetAssetComponentsLoader_LoadAssetComponents_Call {
	return &AssetAssetComponentsLoader_LoadAssetComponents_Call{Call: _e.mock.On("LoadAssetComponents", tx, _a1, version)}
}

func (_c *AssetAssetComponentsLoader_LoadAssetComponents_Call) Run(run func(tx *gorm.DB, _a1 models.Asset, version string)) *AssetAssetComponentsLoader_LoadAssetComponents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(models.Asset), args[2].(string))
	})
	return _c
}

func (_c *AssetAssetComponentsLoader_LoadAssetComponents_Call) Return(_a0 []models.ComponentDependency, _a1 error) *AssetAssetComponentsLoader_LoadAssetComponents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AssetAssetComponentsLoader_LoadAssetComponents_Call) RunAndReturn(run func(*gorm.DB, models.Asset, string) ([]models.ComponentDependency, error)) *AssetAssetComponentsLoader_LoadAssetComponents_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssetAssetComponentsLoader creates a new instance of AssetAssetComponentsLoader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssetAssetComponentsLoader(t interface {
	mock.TestingT
	Cleanup(func())
}) *AssetAssetComponentsLoader {
	mock := &AssetAssetComponentsLoader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
