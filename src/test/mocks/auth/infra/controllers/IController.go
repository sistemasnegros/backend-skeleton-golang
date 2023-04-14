// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// IController is an autogenerated mock type for the IController type
type IController struct {
	mock.Mock
}

// LoadRouter provides a mock function with given fields: _a0
func (_m *IController) LoadRouter(_a0 fiber.Router) {
	_m.Called(_a0)
}

// Login provides a mock function with given fields: _a0
func (_m *IController) Login(_a0 *fiber.Ctx) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: _a0
func (_m *IController) Register(_a0 *fiber.Ctx) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIController interface {
	mock.TestingT
	Cleanup(func())
}

// NewIController creates a new instance of IController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIController(t mockConstructorTestingTNewIController) *IController {
	mock := &IController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
