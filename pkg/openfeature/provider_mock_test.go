// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/openfeature/provider.go

// Package openfeature is a generated GoMock package.
package openfeature

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeatureProvider is a mock of FeatureProvider interface.
type MockFeatureProvider struct {
	ctrl     *gomock.Controller
	recorder *MockFeatureProviderMockRecorder
}

// MockFeatureProviderMockRecorder is the mock recorder for MockFeatureProvider.
type MockFeatureProviderMockRecorder struct {
	mock *MockFeatureProvider
}

// NewMockFeatureProvider creates a new mock instance.
func NewMockFeatureProvider(ctrl *gomock.Controller) *MockFeatureProvider {
	mock := &MockFeatureProvider{ctrl: ctrl}
	mock.recorder = &MockFeatureProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeatureProvider) EXPECT() *MockFeatureProviderMockRecorder {
	return m.recorder
}

// BooleanEvaluation mocks base method.
func (m *MockFeatureProvider) BooleanEvaluation(flag string, defaultValue bool, evalCtx EvaluationContext, options EvaluationOptions) BoolResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BooleanEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(BoolResolutionDetail)
	return ret0
}

// BooleanEvaluation indicates an expected call of BooleanEvaluation.
func (mr *MockFeatureProviderMockRecorder) BooleanEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BooleanEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).BooleanEvaluation), flag, defaultValue, evalCtx, options)
}

// FloatEvaluation mocks base method.
func (m *MockFeatureProvider) FloatEvaluation(flag string, defaultValue float64, evalCtx EvaluationContext, options EvaluationOptions) FloatResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FloatEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(FloatResolutionDetail)
	return ret0
}

// FloatEvaluation indicates an expected call of FloatEvaluation.
func (mr *MockFeatureProviderMockRecorder) FloatEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FloatEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).FloatEvaluation), flag, defaultValue, evalCtx, options)
}

// Hooks mocks base method.
func (m *MockFeatureProvider) Hooks() []Hook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hooks")
	ret0, _ := ret[0].([]Hook)
	return ret0
}

// Hooks indicates an expected call of Hooks.
func (mr *MockFeatureProviderMockRecorder) Hooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hooks", reflect.TypeOf((*MockFeatureProvider)(nil).Hooks))
}

// IntEvaluation mocks base method.
func (m *MockFeatureProvider) IntEvaluation(flag string, defaultValue int64, evalCtx EvaluationContext, options EvaluationOptions) IntResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(IntResolutionDetail)
	return ret0
}

// IntEvaluation indicates an expected call of IntEvaluation.
func (mr *MockFeatureProviderMockRecorder) IntEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).IntEvaluation), flag, defaultValue, evalCtx, options)
}

// Metadata mocks base method.
func (m *MockFeatureProvider) Metadata() Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(Metadata)
	return ret0
}

// Metadata indicates an expected call of Metadata.
func (mr *MockFeatureProviderMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockFeatureProvider)(nil).Metadata))
}

// ObjectEvaluation mocks base method.
func (m *MockFeatureProvider) ObjectEvaluation(flag string, defaultValue interface{}, evalCtx EvaluationContext, options EvaluationOptions) ResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(ResolutionDetail)
	return ret0
}

// ObjectEvaluation indicates an expected call of ObjectEvaluation.
func (mr *MockFeatureProviderMockRecorder) ObjectEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).ObjectEvaluation), flag, defaultValue, evalCtx, options)
}

// StringEvaluation mocks base method.
func (m *MockFeatureProvider) StringEvaluation(flag, defaultValue string, evalCtx EvaluationContext, options EvaluationOptions) StringResolutionDetail {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StringEvaluation", flag, defaultValue, evalCtx, options)
	ret0, _ := ret[0].(StringResolutionDetail)
	return ret0
}

// StringEvaluation indicates an expected call of StringEvaluation.
func (mr *MockFeatureProviderMockRecorder) StringEvaluation(flag, defaultValue, evalCtx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StringEvaluation", reflect.TypeOf((*MockFeatureProvider)(nil).StringEvaluation), flag, defaultValue, evalCtx, options)
}
