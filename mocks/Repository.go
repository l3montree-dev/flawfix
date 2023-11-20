// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository[ID interface{}, T interface{}, Tx interface{}] struct {
	mock.Mock
}

// Create provides a mock function with given fields: tx, t
func (_m *Repository[ID, T, Tx]) Create(tx Tx, t *T) error {
	ret := _m.Called(tx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(Tx, *T) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: tx, id
func (_m *Repository[ID, T, Tx]) Delete(tx Tx, id ID) error {
	ret := _m.Called(tx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(Tx, ID) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ids
func (_m *Repository[ID, T, Tx]) List(ids []ID) ([]T, error) {
	ret := _m.Called(ids)

	var r0 []T
	var r1 error
	if rf, ok := ret.Get(0).(func([]ID) ([]T, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]ID) []T); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]T)
		}
	}

	if rf, ok := ret.Get(1).(func([]ID) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: id
func (_m *Repository[ID, T, Tx]) Read(id ID) (T, error) {
	ret := _m.Called(id)

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(ID) (T, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(ID) T); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(ID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transaction provides a mock function with given fields: _a0
func (_m *Repository[ID, T, Tx]) Transaction(_a0 func(Tx) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(Tx) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: tx, t
func (_m *Repository[ID, T, Tx]) Update(tx Tx, t *T) error {
	ret := _m.Called(tx, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(Tx, *T) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository[ID interface{}, T interface{}, Tx interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository[ID, T, Tx] {
	mock := &Repository[ID, T, Tx]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
