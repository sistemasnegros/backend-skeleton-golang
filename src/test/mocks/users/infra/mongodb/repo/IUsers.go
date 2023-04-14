// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	serviceDomain "backend-skeleton-golang/commons/domain/service"
	usersDomain "backend-skeleton-golang/users/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUsers is an autogenerated mock type for the IUsers type
type IUsers struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *IUsers) Create(_a0 usersDomain.User) (usersDomain.User, error) {
	ret := _m.Called(_a0)

	var r0 usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(usersDomain.User) (usersDomain.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(usersDomain.User) usersDomain.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(usersDomain.User)
	}

	if rf, ok := ret.Get(1).(func(usersDomain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteById provides a mock function with given fields: id
func (_m *IUsers) DeleteById(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: query
func (_m *IUsers) Find(query interface{}) ([]usersDomain.User, error) {
	ret := _m.Called(query)

	var r0 []usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) ([]usersDomain.User, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(interface{}) []usersDomain.User); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]usersDomain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *IUsers) FindById(id string) (*usersDomain.User, error) {
	ret := _m.Called(id)

	var r0 *usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*usersDomain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *usersDomain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersDomain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: query
func (_m *IUsers) FindOne(query interface{}) (*usersDomain.User, error) {
	ret := _m.Called(query)

	var r0 *usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (*usersDomain.User, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(interface{}) *usersDomain.User); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersDomain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPagination provides a mock function with given fields: query, limit, page
func (_m *IUsers) FindPagination(query interface{}, limit int64, page int) (*serviceDomain.PaginationData[usersDomain.User], error) {
	ret := _m.Called(query, limit, page)

	var r0 *serviceDomain.PaginationData[usersDomain.User]
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int64, int) (*serviceDomain.PaginationData[usersDomain.User], error)); ok {
		return rf(query, limit, page)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int64, int) *serviceDomain.PaginationData[usersDomain.User]); ok {
		r0 = rf(query, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*serviceDomain.PaginationData[usersDomain.User])
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int64, int) error); ok {
		r1 = rf(query, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWithNot provides a mock function with given fields: queryNot, query
func (_m *IUsers) FindWithNot(queryNot map[string]interface{}, query map[string]interface{}) (*usersDomain.User, error) {
	ret := _m.Called(queryNot, query)

	var r0 *usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}, map[string]interface{}) (*usersDomain.User, error)); ok {
		return rf(queryNot, query)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, map[string]interface{}) *usersDomain.User); ok {
		r0 = rf(queryNot, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersDomain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}, map[string]interface{}) error); ok {
		r1 = rf(queryNot, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateById provides a mock function with given fields: id, data
func (_m *IUsers) UpdateById(id string, data interface{}) (*usersDomain.User, error) {
	ret := _m.Called(id, data)

	var r0 *usersDomain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}) (*usersDomain.User, error)); ok {
		return rf(id, data)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}) *usersDomain.User); ok {
		r0 = rf(id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usersDomain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}) error); ok {
		r1 = rf(id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIUsers interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUsers creates a new instance of IUsers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUsers(t mockConstructorTestingTNewIUsers) *IUsers {
	mock := &IUsers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
