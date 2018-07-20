// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package store_mocks is a generated GoMock package.
package store_mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	intake_form "go.appointy.com/google/pb/intake_form"
	reflect "reflect"
)

// MockIntakeFormStore is a mock of IntakeFormStore interface
type MockIntakeFormStore struct {
	ctrl     *gomock.Controller
	recorder *MockIntakeFormStoreMockRecorder
}

// MockIntakeFormStoreMockRecorder is the mock recorder for MockIntakeFormStore
type MockIntakeFormStoreMockRecorder struct {
	mock *MockIntakeFormStore
}

// NewMockIntakeFormStore creates a new mock instance
func NewMockIntakeFormStore(ctrl *gomock.Controller) *MockIntakeFormStore {
	mock := &MockIntakeFormStore{ctrl: ctrl}
	mock.recorder = &MockIntakeFormStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIntakeFormStore) EXPECT() *MockIntakeFormStoreMockRecorder {
	return m.recorder
}

// AddIntakeForms mocks base method
func (m *MockIntakeFormStore) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {
	ret := m.ctrl.Call(m, "AddIntakeForms", ctx, in)
	ret0, _ := ret[0].(*intake_form.IntakeFormIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIntakeForms indicates an expected call of AddIntakeForms
func (mr *MockIntakeFormStoreMockRecorder) AddIntakeForms(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIntakeForms", reflect.TypeOf((*MockIntakeFormStore)(nil).AddIntakeForms), ctx, in)
}

// UpdateIntakeForm mocks base method
func (m *MockIntakeFormStore) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm) error {
	ret := m.ctrl.Call(m, "UpdateIntakeForm", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIntakeForm indicates an expected call of UpdateIntakeForm
func (mr *MockIntakeFormStoreMockRecorder) UpdateIntakeForm(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIntakeForm", reflect.TypeOf((*MockIntakeFormStore)(nil).UpdateIntakeForm), ctx, in)
}

// DeleteIntakeForm mocks base method
func (m *MockIntakeFormStore) DeleteIntakeForm(ctx context.Context, ID string) error {
	ret := m.ctrl.Call(m, "DeleteIntakeForm", ctx, ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIntakeForm indicates an expected call of DeleteIntakeForm
func (mr *MockIntakeFormStoreMockRecorder) DeleteIntakeForm(ctx, ID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIntakeForm", reflect.TypeOf((*MockIntakeFormStore)(nil).DeleteIntakeForm), ctx, ID)
}

// GetIntakeForms mocks base method
func (m *MockIntakeFormStore) GetIntakeForms(ctx context.Context, loc string) (*intake_form.IntakeFormList, error) {
	ret := m.ctrl.Call(m, "GetIntakeForms", ctx, loc)
	ret0, _ := ret[0].(*intake_form.IntakeFormList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeForms indicates an expected call of GetIntakeForms
func (mr *MockIntakeFormStoreMockRecorder) GetIntakeForms(ctx, loc interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeForms", reflect.TypeOf((*MockIntakeFormStore)(nil).GetIntakeForms), ctx, loc)
}

// GetIntakeFormById mocks base method
func (m *MockIntakeFormStore) GetIntakeFormById(ctx context.Context, id string) (*intake_form.IntakeForm, error) {
	ret := m.ctrl.Call(m, "GetIntakeFormById", ctx, id)
	ret0, _ := ret[0].(*intake_form.IntakeForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeFormById indicates an expected call of GetIntakeFormById
func (mr *MockIntakeFormStoreMockRecorder) GetIntakeFormById(ctx, id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeFormById", reflect.TypeOf((*MockIntakeFormStore)(nil).GetIntakeFormById), ctx, id)
}
