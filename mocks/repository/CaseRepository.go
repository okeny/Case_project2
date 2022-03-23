// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entity "project-cases/entity"

	mock "github.com/stretchr/testify/mock"
)

// CaseRepository is an autogenerated mock type for the CaseRepository type
type CaseRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *CaseRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *CaseRepository) FindAll() ([]entity.Case, error) {
	ret := _m.Called()

	var r0 []entity.Case
	if rf, ok := ret.Get(0).(func() []entity.Case); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Case)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *CaseRepository) FindById(id int) (*entity.Case, error) {
	ret := _m.Called(id)

	var r0 *entity.Case
	if rf, ok := ret.Get(0).(func(int) *entity.Case); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Case)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: Case
func (_m *CaseRepository) Save(Case *entity.Case) (*entity.Case, error) {
	ret := _m.Called(Case)

	var r0 *entity.Case
	if rf, ok := ret.Get(0).(func(*entity.Case) *entity.Case); ok {
		r0 = rf(Case)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Case)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Case) error); ok {
		r1 = rf(Case)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCase provides a mock function with given fields: c
func (_m *CaseRepository) UpdateCase(c *entity.Case) (*entity.Case, error) {
	ret := _m.Called(c)

	var r0 *entity.Case
	if rf, ok := ret.Get(0).(func(*entity.Case) *entity.Case); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Case)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Case) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
