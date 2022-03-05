// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sarchlab/zeonica/cgra (interfaces: Device,Tile)

package api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cgra "github.com/sarchlab/zeonica/cgra"
	sim "gitlab.com/akita/akita/v2/sim"
)

// MockDevice is a mock of Device interface.
type MockDevice struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceMockRecorder
}

// MockDeviceMockRecorder is the mock recorder for MockDevice.
type MockDeviceMockRecorder struct {
	mock *MockDevice
}

// NewMockDevice creates a new mock instance.
func NewMockDevice(ctrl *gomock.Controller) *MockDevice {
	mock := &MockDevice{ctrl: ctrl}
	mock.recorder = &MockDeviceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevice) EXPECT() *MockDeviceMockRecorder {
	return m.recorder
}

// GetSidePorts mocks base method.
func (m *MockDevice) GetSidePorts(arg0 cgra.Side, arg1 [2]int) []sim.Port {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSidePorts", arg0, arg1)
	ret0, _ := ret[0].([]sim.Port)
	return ret0
}

// GetSidePorts indicates an expected call of GetSidePorts.
func (mr *MockDeviceMockRecorder) GetSidePorts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSidePorts", reflect.TypeOf((*MockDevice)(nil).GetSidePorts), arg0, arg1)
}

// GetSize mocks base method.
func (m *MockDevice) GetSize() (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSize")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetSize indicates an expected call of GetSize.
func (mr *MockDeviceMockRecorder) GetSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSize", reflect.TypeOf((*MockDevice)(nil).GetSize))
}

// GetTile mocks base method.
func (m *MockDevice) GetTile(arg0, arg1 int) cgra.Tile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTile", arg0, arg1)
	ret0, _ := ret[0].(cgra.Tile)
	return ret0
}

// GetTile indicates an expected call of GetTile.
func (mr *MockDeviceMockRecorder) GetTile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTile", reflect.TypeOf((*MockDevice)(nil).GetTile), arg0, arg1)
}

// MockTile is a mock of Tile interface.
type MockTile struct {
	ctrl     *gomock.Controller
	recorder *MockTileMockRecorder
}

// MockTileMockRecorder is the mock recorder for MockTile.
type MockTileMockRecorder struct {
	mock *MockTile
}

// NewMockTile creates a new mock instance.
func NewMockTile(ctrl *gomock.Controller) *MockTile {
	mock := &MockTile{ctrl: ctrl}
	mock.recorder = &MockTileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTile) EXPECT() *MockTileMockRecorder {
	return m.recorder
}

// GetPort mocks base method.
func (m *MockTile) GetPort(arg0 cgra.Side) sim.Port {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", arg0)
	ret0, _ := ret[0].(sim.Port)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockTileMockRecorder) GetPort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockTile)(nil).GetPort), arg0)
}

// MapProgram mocks base method.
func (m *MockTile) MapProgram(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MapProgram", arg0)
}

// MapProgram indicates an expected call of MapProgram.
func (mr *MockTileMockRecorder) MapProgram(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapProgram", reflect.TypeOf((*MockTile)(nil).MapProgram), arg0)
}

// SetRemotePort mocks base method.
func (m *MockTile) SetRemotePort(arg0 cgra.Side, arg1 sim.Port) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRemotePort", arg0, arg1)
}

// SetRemotePort indicates an expected call of SetRemotePort.
func (mr *MockTileMockRecorder) SetRemotePort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRemotePort", reflect.TypeOf((*MockTile)(nil).SetRemotePort), arg0, arg1)
}
