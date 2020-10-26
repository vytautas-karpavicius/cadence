// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: mutable_state_task_generator.go

// Package execution is a generated GoMock package.
package execution

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	shared "github.com/uber/cadence/.gen/go/shared"
)

// MockMutableStateTaskGenerator is a mock of MutableStateTaskGenerator interface
type MockMutableStateTaskGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockMutableStateTaskGeneratorMockRecorder
}

// MockMutableStateTaskGeneratorMockRecorder is the mock recorder for MockMutableStateTaskGenerator
type MockMutableStateTaskGeneratorMockRecorder struct {
	mock *MockMutableStateTaskGenerator
}

// NewMockMutableStateTaskGenerator creates a new mock instance
func NewMockMutableStateTaskGenerator(ctrl *gomock.Controller) *MockMutableStateTaskGenerator {
	mock := &MockMutableStateTaskGenerator{ctrl: ctrl}
	mock.recorder = &MockMutableStateTaskGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMutableStateTaskGenerator) EXPECT() *MockMutableStateTaskGeneratorMockRecorder {
	return m.recorder
}

// GenerateWorkflowStartTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateWorkflowStartTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowStartTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowStartTasks indicates an expected call of GenerateWorkflowStartTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateWorkflowStartTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowStartTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateWorkflowStartTasks), now, startEvent)
}

// GenerateWorkflowCloseTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateWorkflowCloseTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowCloseTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowCloseTasks indicates an expected call of GenerateWorkflowCloseTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateWorkflowCloseTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowCloseTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateWorkflowCloseTasks), now)
}

// GenerateRecordWorkflowStartedTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateRecordWorkflowStartedTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRecordWorkflowStartedTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateRecordWorkflowStartedTasks indicates an expected call of GenerateRecordWorkflowStartedTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateRecordWorkflowStartedTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRecordWorkflowStartedTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateRecordWorkflowStartedTasks), now, startEvent)
}

// GenerateDelayedDecisionTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateDelayedDecisionTasks(now time.Time, startEvent *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDelayedDecisionTasks", now, startEvent)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateDelayedDecisionTasks indicates an expected call of GenerateDelayedDecisionTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateDelayedDecisionTasks(now, startEvent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDelayedDecisionTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateDelayedDecisionTasks), now, startEvent)
}

// GenerateDecisionScheduleTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateDecisionScheduleTasks(now time.Time, decisionScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDecisionScheduleTasks", now, decisionScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateDecisionScheduleTasks indicates an expected call of GenerateDecisionScheduleTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateDecisionScheduleTasks(now, decisionScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDecisionScheduleTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateDecisionScheduleTasks), now, decisionScheduleID)
}

// GenerateDecisionStartTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateDecisionStartTasks(now time.Time, decisionScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateDecisionStartTasks", now, decisionScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateDecisionStartTasks indicates an expected call of GenerateDecisionStartTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateDecisionStartTasks(now, decisionScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateDecisionStartTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateDecisionStartTasks), now, decisionScheduleID)
}

// GenerateActivityTransferTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateActivityTransferTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityTransferTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityTransferTasks indicates an expected call of GenerateActivityTransferTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateActivityTransferTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityTransferTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateActivityTransferTasks), now, event)
}

// GenerateActivityRetryTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateActivityRetryTasks(activityScheduleID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityRetryTasks", activityScheduleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityRetryTasks indicates an expected call of GenerateActivityRetryTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateActivityRetryTasks(activityScheduleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityRetryTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateActivityRetryTasks), activityScheduleID)
}

// GenerateChildWorkflowTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateChildWorkflowTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateChildWorkflowTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateChildWorkflowTasks indicates an expected call of GenerateChildWorkflowTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateChildWorkflowTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateChildWorkflowTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateChildWorkflowTasks), now, event)
}

// GenerateRequestCancelExternalTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateRequestCancelExternalTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateRequestCancelExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateRequestCancelExternalTasks indicates an expected call of GenerateRequestCancelExternalTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateRequestCancelExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateRequestCancelExternalTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateRequestCancelExternalTasks), now, event)
}

// GenerateSignalExternalTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateSignalExternalTasks(now time.Time, event *shared.HistoryEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSignalExternalTasks", now, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateSignalExternalTasks indicates an expected call of GenerateSignalExternalTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateSignalExternalTasks(now, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSignalExternalTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateSignalExternalTasks), now, event)
}

// GenerateWorkflowSearchAttrTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateWorkflowSearchAttrTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowSearchAttrTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowSearchAttrTasks indicates an expected call of GenerateWorkflowSearchAttrTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateWorkflowSearchAttrTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowSearchAttrTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateWorkflowSearchAttrTasks), now)
}

// GenerateWorkflowResetTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateWorkflowResetTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateWorkflowResetTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateWorkflowResetTasks indicates an expected call of GenerateWorkflowResetTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateWorkflowResetTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateWorkflowResetTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateWorkflowResetTasks), now)
}

// GenerateActivityTimerTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateActivityTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateActivityTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateActivityTimerTasks indicates an expected call of GenerateActivityTimerTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateActivityTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateActivityTimerTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateActivityTimerTasks), now)
}

// GenerateUserTimerTasks mocks base method
func (m *MockMutableStateTaskGenerator) GenerateUserTimerTasks(now time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUserTimerTasks", now)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateUserTimerTasks indicates an expected call of GenerateUserTimerTasks
func (mr *MockMutableStateTaskGeneratorMockRecorder) GenerateUserTimerTasks(now interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUserTimerTasks", reflect.TypeOf((*MockMutableStateTaskGenerator)(nil).GenerateUserTimerTasks), now)
}
