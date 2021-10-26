// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	gc "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/gc"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockGcClientService is an autogenerated mock type for the ClientService type
type MockGcClientService struct {
	mock.Mock
}

// CreateGCSchedule provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) CreateGCSchedule(params *gc.CreateGCScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*gc.CreateGCScheduleCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.CreateGCScheduleCreated
	if rf, ok := ret.Get(0).(func(*gc.CreateGCScheduleParams, runtime.ClientAuthInfoWriter) *gc.CreateGCScheduleCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.CreateGCScheduleCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.CreateGCScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGC provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) GetGC(params *gc.GetGCParams, authInfo runtime.ClientAuthInfoWriter) (*gc.GetGCOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.GetGCOK
	if rf, ok := ret.Get(0).(func(*gc.GetGCParams, runtime.ClientAuthInfoWriter) *gc.GetGCOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.GetGCOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.GetGCParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGCHistory provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) GetGCHistory(params *gc.GetGCHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*gc.GetGCHistoryOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.GetGCHistoryOK
	if rf, ok := ret.Get(0).(func(*gc.GetGCHistoryParams, runtime.ClientAuthInfoWriter) *gc.GetGCHistoryOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.GetGCHistoryOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.GetGCHistoryParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGCLog provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) GetGCLog(params *gc.GetGCLogParams, authInfo runtime.ClientAuthInfoWriter) (*gc.GetGCLogOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.GetGCLogOK
	if rf, ok := ret.Get(0).(func(*gc.GetGCLogParams, runtime.ClientAuthInfoWriter) *gc.GetGCLogOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.GetGCLogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.GetGCLogParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGCSchedule provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) GetGCSchedule(params *gc.GetGCScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*gc.GetGCScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.GetGCScheduleOK
	if rf, ok := ret.Get(0).(func(*gc.GetGCScheduleParams, runtime.ClientAuthInfoWriter) *gc.GetGCScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.GetGCScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.GetGCScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockGcClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateGCSchedule provides a mock function with given fields: params, authInfo
func (_m *MockGcClientService) UpdateGCSchedule(params *gc.UpdateGCScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*gc.UpdateGCScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *gc.UpdateGCScheduleOK
	if rf, ok := ret.Get(0).(func(*gc.UpdateGCScheduleParams, runtime.ClientAuthInfoWriter) *gc.UpdateGCScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gc.UpdateGCScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gc.UpdateGCScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
