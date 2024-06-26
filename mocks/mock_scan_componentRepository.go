// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// ScanComponentRepository is an autogenerated mock type for the componentRepository type
type ScanComponentRepository struct {
	mock.Mock
}

type ScanComponentRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ScanComponentRepository) EXPECT() *ScanComponentRepository_Expecter {
	return &ScanComponentRepository_Expecter{mock: &_m.Mock}
}

// SaveBatch provides a mock function with given fields: tx, components
func (_m *ScanComponentRepository) SaveBatch(tx *gorm.DB, components []models.Component) error {
	ret := _m.Called(tx, components)

	if len(ret) == 0 {
		panic("no return value specified for SaveBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.Component) error); ok {
		r0 = rf(tx, components)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ScanComponentRepository_SaveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBatch'
type ScanComponentRepository_SaveBatch_Call struct {
	*mock.Call
}

// SaveBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - components []models.Component
func (_e *ScanComponentRepository_Expecter) SaveBatch(tx interface{}, components interface{}) *ScanComponentRepository_SaveBatch_Call {
	return &ScanComponentRepository_SaveBatch_Call{Call: _e.mock.On("SaveBatch", tx, components)}
}

func (_c *ScanComponentRepository_SaveBatch_Call) Run(run func(tx *gorm.DB, components []models.Component)) *ScanComponentRepository_SaveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.Component))
	})
	return _c
}

func (_c *ScanComponentRepository_SaveBatch_Call) Return(_a0 error) *ScanComponentRepository_SaveBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ScanComponentRepository_SaveBatch_Call) RunAndReturn(run func(*gorm.DB, []models.Component) error) *ScanComponentRepository_SaveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// NewScanComponentRepository creates a new instance of ScanComponentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScanComponentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScanComponentRepository {
	mock := &ScanComponentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
