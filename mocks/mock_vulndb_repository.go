// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	core "github.com/l3montree-dev/devguard/internal/core"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/l3montree-dev/devguard/internal/database/models"
)

// VulndbRepository is an autogenerated mock type for the repository type
type VulndbRepository struct {
	mock.Mock
}

type VulndbRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *VulndbRepository) EXPECT() *VulndbRepository_Expecter {
	return &VulndbRepository_Expecter{mock: &_m.Mock}
}

// FindAllListPaged provides a mock function with given fields: tx, pageInfo, filter, sort
func (_m *VulndbRepository) FindAllListPaged(tx *gorm.DB, pageInfo core.PageInfo, filter []core.FilterQuery, sort []core.SortQuery) (core.Paged[models.CVE], error) {
	ret := _m.Called(tx, pageInfo, filter, sort)

	if len(ret) == 0 {
		panic("no return value specified for FindAllListPaged")
	}

	var r0 core.Paged[models.CVE]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, core.PageInfo, []core.FilterQuery, []core.SortQuery) (core.Paged[models.CVE], error)); ok {
		return rf(tx, pageInfo, filter, sort)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, core.PageInfo, []core.FilterQuery, []core.SortQuery) core.Paged[models.CVE]); ok {
		r0 = rf(tx, pageInfo, filter, sort)
	} else {
		r0 = ret.Get(0).(core.Paged[models.CVE])
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, core.PageInfo, []core.FilterQuery, []core.SortQuery) error); ok {
		r1 = rf(tx, pageInfo, filter, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbRepository_FindAllListPaged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllListPaged'
type VulndbRepository_FindAllListPaged_Call struct {
	*mock.Call
}

// FindAllListPaged is a helper method to define mock.On call
//   - tx *gorm.DB
//   - pageInfo core.PageInfo
//   - filter []core.FilterQuery
//   - sort []core.SortQuery
func (_e *VulndbRepository_Expecter) FindAllListPaged(tx interface{}, pageInfo interface{}, filter interface{}, sort interface{}) *VulndbRepository_FindAllListPaged_Call {
	return &VulndbRepository_FindAllListPaged_Call{Call: _e.mock.On("FindAllListPaged", tx, pageInfo, filter, sort)}
}

func (_c *VulndbRepository_FindAllListPaged_Call) Run(run func(tx *gorm.DB, pageInfo core.PageInfo, filter []core.FilterQuery, sort []core.SortQuery)) *VulndbRepository_FindAllListPaged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(core.PageInfo), args[2].([]core.FilterQuery), args[3].([]core.SortQuery))
	})
	return _c
}

func (_c *VulndbRepository_FindAllListPaged_Call) Return(_a0 core.Paged[models.CVE], _a1 error) *VulndbRepository_FindAllListPaged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbRepository_FindAllListPaged_Call) RunAndReturn(run func(*gorm.DB, core.PageInfo, []core.FilterQuery, []core.SortQuery) (core.Paged[models.CVE], error)) *VulndbRepository_FindAllListPaged_Call {
	_c.Call.Return(run)
	return _c
}

// FindCVE provides a mock function with given fields: tx, cveId
func (_m *VulndbRepository) FindCVE(tx *gorm.DB, cveId string) (models.CVE, error) {
	ret := _m.Called(tx, cveId)

	if len(ret) == 0 {
		panic("no return value specified for FindCVE")
	}

	var r0 models.CVE
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (models.CVE, error)); ok {
		return rf(tx, cveId)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) models.CVE); ok {
		r0 = rf(tx, cveId)
	} else {
		r0 = ret.Get(0).(models.CVE)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(tx, cveId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VulndbRepository_FindCVE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindCVE'
type VulndbRepository_FindCVE_Call struct {
	*mock.Call
}

// FindCVE is a helper method to define mock.On call
//   - tx *gorm.DB
//   - cveId string
func (_e *VulndbRepository_Expecter) FindCVE(tx interface{}, cveId interface{}) *VulndbRepository_FindCVE_Call {
	return &VulndbRepository_FindCVE_Call{Call: _e.mock.On("FindCVE", tx, cveId)}
}

func (_c *VulndbRepository_FindCVE_Call) Run(run func(tx *gorm.DB, cveId string)) *VulndbRepository_FindCVE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *VulndbRepository_FindCVE_Call) Return(_a0 models.CVE, _a1 error) *VulndbRepository_FindCVE_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VulndbRepository_FindCVE_Call) RunAndReturn(run func(*gorm.DB, string) (models.CVE, error)) *VulndbRepository_FindCVE_Call {
	_c.Call.Return(run)
	return _c
}

// NewVulndbRepository creates a new instance of VulndbRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVulndbRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *VulndbRepository {
	mock := &VulndbRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
