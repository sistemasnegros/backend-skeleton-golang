// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	authDTO "backend-skeleton-golang/auth/app/dto"

	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// Login provides a mock function with given fields: body
func (_m *IService) Login(body *authDTO.Login) (int, interface{}) {
	ret := _m.Called(body)

	var r0 int
	var r1 interface{}
	if rf, ok := ret.Get(0).(func(*authDTO.Login) (int, interface{})); ok {
		return rf(body)
	}
	if rf, ok := ret.Get(0).(func(*authDTO.Login) int); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*authDTO.Login) interface{}); ok {
		r1 = rf(body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: body
func (_m *IService) Register(body *authDTO.Register) (int, interface{}) {
	ret := _m.Called(body)

	var r0 int
	var r1 interface{}
	if rf, ok := ret.Get(0).(func(*authDTO.Register) (int, interface{})); ok {
		return rf(body)
	}
	if rf, ok := ret.Get(0).(func(*authDTO.Register) int); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*authDTO.Register) interface{}); ok {
		r1 = rf(body)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewIService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIService(t mockConstructorTestingTNewIService) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
