// Code generated by mockery v2.2.1. DO NOT EDIT.

package validators

import mock "github.com/stretchr/testify/mock"

// MockValidator is an autogenerated mock type for the Validator type
type MockValidator struct {
	mock.Mock
}

// Struct provides a mock function with given fields: current
func (_m *MockValidator) Struct(current interface{}) error {
	ret := _m.Called(current)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(current)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}