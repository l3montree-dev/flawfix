// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/devguard/internal/database/models"
	mock "github.com/stretchr/testify/mock"
)

// AssetAssetService is an autogenerated mock type for the assetService type
type AssetAssetService struct {
	mock.Mock
}

type AssetAssetService_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetAssetService) EXPECT() *AssetAssetService_Expecter {
	return &AssetAssetService_Expecter{mock: &_m.Mock}
}

// UpdateAssetRequirements provides a mock function with given fields: _a0, responsible, justification
func (_m *AssetAssetService) UpdateAssetRequirements(_a0 models.Asset, responsible string, justification string) error {
	ret := _m.Called(_a0, responsible, justification)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAssetRequirements")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Asset, string, string) error); ok {
		r0 = rf(_a0, responsible, justification)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetAssetService_UpdateAssetRequirements_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAssetRequirements'
type AssetAssetService_UpdateAssetRequirements_Call struct {
	*mock.Call
}

// UpdateAssetRequirements is a helper method to define mock.On call
//   - _a0 models.Asset
//   - responsible string
//   - justification string
func (_e *AssetAssetService_Expecter) UpdateAssetRequirements(_a0 interface{}, responsible interface{}, justification interface{}) *AssetAssetService_UpdateAssetRequirements_Call {
	return &AssetAssetService_UpdateAssetRequirements_Call{Call: _e.mock.On("UpdateAssetRequirements", _a0, responsible, justification)}
}

func (_c *AssetAssetService_UpdateAssetRequirements_Call) Run(run func(_a0 models.Asset, responsible string, justification string)) *AssetAssetService_UpdateAssetRequirements_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Asset), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *AssetAssetService_UpdateAssetRequirements_Call) Return(_a0 error) *AssetAssetService_UpdateAssetRequirements_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetAssetService_UpdateAssetRequirements_Call) RunAndReturn(run func(models.Asset, string, string) error) *AssetAssetService_UpdateAssetRequirements_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssetAssetService creates a new instance of AssetAssetService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssetAssetService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AssetAssetService {
	mock := &AssetAssetService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
