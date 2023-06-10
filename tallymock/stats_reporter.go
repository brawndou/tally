// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber-go/tally (interfaces: StatsReporter)

// Package tallymock is a generated GoMock package.
package tallymock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	tally "github.com/uber-go/tally/v4"
)

// MockStatsReporter is a mock of StatsReporter interface.
type MockStatsReporter struct {
	ctrl     *gomock.Controller
	recorder *MockStatsReporterMockRecorder
}

// MockStatsReporterMockRecorder is the mock recorder for MockStatsReporter.
type MockStatsReporterMockRecorder struct {
	mock *MockStatsReporter
}

// NewMockStatsReporter creates a new mock instance.
func NewMockStatsReporter(ctrl *gomock.Controller) *MockStatsReporter {
	mock := &MockStatsReporter{ctrl: ctrl}
	mock.recorder = &MockStatsReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsReporter) EXPECT() *MockStatsReporterMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockStatsReporter) Capabilities() tally.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(tally.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockStatsReporterMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockStatsReporter)(nil).Capabilities))
}

// Flush mocks base method.
func (m *MockStatsReporter) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush.
func (mr *MockStatsReporterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockStatsReporter)(nil).Flush))
}

// ReportCounter mocks base method.
func (m *MockStatsReporter) ReportCounter(arg0 string, arg1 map[string]string, arg2 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportCounter", arg0, arg1, arg2)
}

// ReportCounter indicates an expected call of ReportCounter.
func (mr *MockStatsReporterMockRecorder) ReportCounter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportCounter", reflect.TypeOf((*MockStatsReporter)(nil).ReportCounter), arg0, arg1, arg2)
}

// ReportGauge mocks base method.
func (m *MockStatsReporter) ReportGauge(arg0 string, arg1 map[string]string, arg2 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportGauge", arg0, arg1, arg2)
}

// ReportGauge indicates an expected call of ReportGauge.
func (mr *MockStatsReporterMockRecorder) ReportGauge(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportGauge", reflect.TypeOf((*MockStatsReporter)(nil).ReportGauge), arg0, arg1, arg2)
}

// ReportHistogramDurationSamples mocks base method.
func (m *MockStatsReporter) ReportHistogramDurationSamples(arg0 string, arg1 map[string]string, arg2 tally.Buckets, arg3, arg4 time.Duration, arg5 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportHistogramDurationSamples", arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReportHistogramDurationSamples indicates an expected call of ReportHistogramDurationSamples.
func (mr *MockStatsReporterMockRecorder) ReportHistogramDurationSamples(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportHistogramDurationSamples", reflect.TypeOf((*MockStatsReporter)(nil).ReportHistogramDurationSamples), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReportHistogramValueSamples mocks base method.
func (m *MockStatsReporter) ReportHistogramValueSamples(arg0 string, arg1 map[string]string, arg2 tally.Buckets, arg3, arg4 float64, arg5 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportHistogramValueSamples", arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReportHistogramValueSamples indicates an expected call of ReportHistogramValueSamples.
func (mr *MockStatsReporterMockRecorder) ReportHistogramValueSamples(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportHistogramValueSamples", reflect.TypeOf((*MockStatsReporter)(nil).ReportHistogramValueSamples), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ReportTimer mocks base method.
func (m *MockStatsReporter) ReportTimer(arg0 string, arg1 map[string]string, arg2 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReportTimer", arg0, arg1, arg2)
}

// ReportTimer indicates an expected call of ReportTimer.
func (mr *MockStatsReporterMockRecorder) ReportTimer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportTimer", reflect.TypeOf((*MockStatsReporter)(nil).ReportTimer), arg0, arg1, arg2)
}
