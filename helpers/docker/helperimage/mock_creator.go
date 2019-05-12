// Code generated by mockery v1.0.0. DO NOT EDIT.

// This comment works around https://github.com/vektra/mockery/issues/155

package helperimage

import mock "github.com/stretchr/testify/mock"

// mockCreator is an autogenerated mock type for the creator type
type mockCreator struct {
	mock.Mock
}

// Create provides a mock function with given fields: revision, cfg
func (_m *mockCreator) Create(revision string, cfg Config) (Info, error) {
	ret := _m.Called(revision, cfg)

	var r0 Info
	if rf, ok := ret.Get(0).(func(string, Config) Info); ok {
		r0 = rf(revision, cfg)
	} else {
		r0 = ret.Get(0).(Info)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, Config) error); ok {
		r1 = rf(revision, cfg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
