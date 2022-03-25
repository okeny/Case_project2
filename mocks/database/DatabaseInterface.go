// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// DatabaseInterface is an autogenerated mock type for the DatabaseInterface type
type DatabaseInterface struct {
	mock.Mock
}

// GetConnection provides a mock function with given fields:
func (_m *DatabaseInterface) GetConnection() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}
