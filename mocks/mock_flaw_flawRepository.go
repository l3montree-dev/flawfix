// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "github.com/l3montree-dev/flawfix/internal/database/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// FlawFlawRepository is an autogenerated mock type for the flawRepository type
type FlawFlawRepository struct {
	mock.Mock
}

type FlawFlawRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *FlawFlawRepository) EXPECT() *FlawFlawRepository_Expecter {
	return &FlawFlawRepository_Expecter{mock: &_m.Mock}
}

// SaveBatch provides a mock function with given fields: db, flaws
func (_m *FlawFlawRepository) SaveBatch(db *gorm.DB, flaws []models.Flaw) error {
	ret := _m.Called(db, flaws)

	if len(ret) == 0 {
		panic("no return value specified for SaveBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.Flaw) error); ok {
		r0 = rf(db, flaws)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlawFlawRepository_SaveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBatch'
type FlawFlawRepository_SaveBatch_Call struct {
	*mock.Call
}

// SaveBatch is a helper method to define mock.On call
//   - db *gorm.DB
//   - flaws []models.Flaw
func (_e *FlawFlawRepository_Expecter) SaveBatch(db interface{}, flaws interface{}) *FlawFlawRepository_SaveBatch_Call {
	return &FlawFlawRepository_SaveBatch_Call{Call: _e.mock.On("SaveBatch", db, flaws)}
}

func (_c *FlawFlawRepository_SaveBatch_Call) Run(run func(db *gorm.DB, flaws []models.Flaw)) *FlawFlawRepository_SaveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.Flaw))
	})
	return _c
}

func (_c *FlawFlawRepository_SaveBatch_Call) Return(_a0 error) *FlawFlawRepository_SaveBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlawFlawRepository_SaveBatch_Call) RunAndReturn(run func(*gorm.DB, []models.Flaw) error) *FlawFlawRepository_SaveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: txFunc
func (_m *FlawFlawRepository) Transaction(txFunc func(*gorm.DB) error) error {
	ret := _m.Called(txFunc)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(txFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FlawFlawRepository_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type FlawFlawRepository_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - txFunc func(*gorm.DB) error
func (_e *FlawFlawRepository_Expecter) Transaction(txFunc interface{}) *FlawFlawRepository_Transaction_Call {
	return &FlawFlawRepository_Transaction_Call{Call: _e.mock.On("Transaction", txFunc)}
}

func (_c *FlawFlawRepository_Transaction_Call) Run(run func(txFunc func(*gorm.DB) error)) *FlawFlawRepository_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *FlawFlawRepository_Transaction_Call) Return(_a0 error) *FlawFlawRepository_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FlawFlawRepository_Transaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *FlawFlawRepository_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewFlawFlawRepository creates a new instance of FlawFlawRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlawFlawRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlawFlawRepository {
	mock := &FlawFlawRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
