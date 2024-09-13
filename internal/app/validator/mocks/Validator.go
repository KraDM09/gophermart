// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Validator is an autogenerated mock type for the Validator type
type Validator struct {
	mock.Mock
}

// Initialize provides a mock function with given fields:
func (_m *Validator) Initialize() {
	_m.Called()
}

// Struct provides a mock function with given fields: s
func (_m *Validator) Struct(s interface{}) error {
	ret := _m.Called(s)

	if len(ret) == 0 {
		panic("no return value specified for Struct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewValidator creates a new instance of Validator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidator(t interface {
	mock.TestingT
	Cleanup(func())
}) *Validator {
	mock := &Validator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}