// Code generated by mockery v1.0.0. DO NOT EDIT.

package container

import context "context"
import mock "github.com/stretchr/testify/mock"
import net "net"
import time "time"

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// KillContainer provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockClient) KillContainer(_a0 context.Context, _a1 Container, _a2 string, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, string, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListContainers provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) ListContainers(_a0 context.Context, _a1 FilterFunc, _a2 ListOpts) ([]Container, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []Container
	if rf, ok := ret.Get(0).(func(context.Context, FilterFunc, ListOpts) []Container); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, FilterFunc, ListOpts) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetemContainer provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8
func (_m *MockClient) NetemContainer(_a0 context.Context, _a1 Container, _a2 string, _a3 []string, _a4 []*net.IPNet, _a5 time.Duration, _a6 string, _a7 bool, _a8 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, string, []string, []*net.IPNet, time.Duration, string, bool, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6, _a7, _a8)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PauseContainer provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) PauseContainer(_a0 context.Context, _a1 Container, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveContainer provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *MockClient) RemoveContainer(_a0 context.Context, _a1 Container, _a2 bool, _a3 bool, _a4 bool, _a5 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, bool, bool, bool, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartContainer provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) StartContainer(_a0 context.Context, _a1 Container, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopContainer provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockClient) StopContainer(_a0 context.Context, _a1 Container, _a2 int, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopContainerWithID provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *MockClient) StopContainerWithID(_a0 context.Context, _a1 string, _a2 time.Duration, _a3 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopNetemContainer provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *MockClient) StopNetemContainer(_a0 context.Context, _a1 Container, _a2 string, _a3 []*net.IPNet, _a4 string, _a5 bool, _a6 bool) error {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, string, []*net.IPNet, string, bool, bool) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StressContainer provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5, _a6
func (_m *MockClient) StressContainer(_a0 context.Context, _a1 Container, _a2 []string, _a3 string, _a4 bool, _a5 time.Duration, _a6 bool) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5, _a6)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, Container, []string, string, bool, time.Duration, bool) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, Container, []string, string, bool, time.Duration, bool) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5, _a6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnpauseContainer provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) UnpauseContainer(_a0 context.Context, _a1 Container, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Container, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
