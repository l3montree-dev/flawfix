// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// AssetFlawService is an autogenerated mock type for the flawService type
type AssetFlawService struct {
	mock.Mock
}

type AssetFlawService_Expecter struct {
	mock *mock.Mock
}

func (_m *AssetFlawService) EXPECT() *AssetFlawService_Expecter {
	return &AssetFlawService_Expecter{mock: &_m.Mock}
}

// UserDetectedFlaws provides a mock function with given fields: tx, userID, flaws
func (_m *AssetFlawService) UserDetectedFlaws(tx *gorm.DB, userID string, flaws []models.Flaw) error {
	ret := _m.Called(tx, userID, flaws)

	if len(ret) == 0 {
		panic("no return value specified for UserDetectedFlaws")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, []models.Flaw) error); ok {
		r0 = rf(tx, userID, flaws)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetFlawService_UserDetectedFlaws_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserDetectedFlaws'
type AssetFlawService_UserDetectedFlaws_Call struct {
	*mock.Call
}

// UserDetectedFlaws is a helper method to define mock.On call
//   - tx *gorm.DB
//   - userID string
//   - flaws []models.Flaw
func (_e *AssetFlawService_Expecter) UserDetectedFlaws(tx interface{}, userID interface{}, flaws interface{}) *AssetFlawService_UserDetectedFlaws_Call {
	return &AssetFlawService_UserDetectedFlaws_Call{Call: _e.mock.On("UserDetectedFlaws", tx, userID, flaws)}
}

func (_c *AssetFlawService_UserDetectedFlaws_Call) Run(run func(tx *gorm.DB, userID string, flaws []models.Flaw)) *AssetFlawService_UserDetectedFlaws_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string), args[2].([]models.Flaw))
	})
	return _c
}

func (_c *AssetFlawService_UserDetectedFlaws_Call) Return(_a0 error) *AssetFlawService_UserDetectedFlaws_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetFlawService_UserDetectedFlaws_Call) RunAndReturn(run func(*gorm.DB, string, []models.Flaw) error) *AssetFlawService_UserDetectedFlaws_Call {
	_c.Call.Return(run)
	return _c
}

// UserFixedFlaws provides a mock function with given fields: tx, userID, flaws
func (_m *AssetFlawService) UserFixedFlaws(tx *gorm.DB, userID string, flaws []models.Flaw) error {
	ret := _m.Called(tx, userID, flaws)

	if len(ret) == 0 {
		panic("no return value specified for UserFixedFlaws")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, []models.Flaw) error); ok {
		r0 = rf(tx, userID, flaws)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssetFlawService_UserFixedFlaws_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserFixedFlaws'
type AssetFlawService_UserFixedFlaws_Call struct {
	*mock.Call
}

// UserFixedFlaws is a helper method to define mock.On call
//   - tx *gorm.DB
//   - userID string
//   - flaws []models.Flaw
func (_e *AssetFlawService_Expecter) UserFixedFlaws(tx interface{}, userID interface{}, flaws interface{}) *AssetFlawService_UserFixedFlaws_Call {
	return &AssetFlawService_UserFixedFlaws_Call{Call: _e.mock.On("UserFixedFlaws", tx, userID, flaws)}
}

func (_c *AssetFlawService_UserFixedFlaws_Call) Run(run func(tx *gorm.DB, userID string, flaws []models.Flaw)) *AssetFlawService_UserFixedFlaws_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string), args[2].([]models.Flaw))
	})
	return _c
}

func (_c *AssetFlawService_UserFixedFlaws_Call) Return(_a0 error) *AssetFlawService_UserFixedFlaws_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AssetFlawService_UserFixedFlaws_Call) RunAndReturn(run func(*gorm.DB, string, []models.Flaw) error) *AssetFlawService_UserFixedFlaws_Call {
	_c.Call.Return(run)
	return _c
}

// NewAssetFlawService creates a new instance of AssetFlawService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAssetFlawService(t interface {
	mock.TestingT
	Cleanup(func())
}) *AssetFlawService {
	mock := &AssetFlawService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
