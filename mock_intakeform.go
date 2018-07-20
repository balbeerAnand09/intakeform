// Code generated by MockGen. DO NOT EDIT.
// Source: ..\pb\intake_form\intake_form.pb.go

// Package servers_mocks is a generated GoMock package.
package servers_mocks

import (
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	intake_form "go.appointy.com/google/pb/intake_form"
	locations "go.appointy.com/google/pb/locations"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockIntakeFormsClient is a mock of IntakeFormsClient interface
type MockIntakeFormsClient struct {
	ctrl     *gomock.Controller
	recorder *MockIntakeFormsClientMockRecorder
}

// MockIntakeFormsClientMockRecorder is the mock recorder for MockIntakeFormsClient
type MockIntakeFormsClientMockRecorder struct {
	mock *MockIntakeFormsClient
}

// NewMockIntakeFormsClient creates a new mock instance
func NewMockIntakeFormsClient(ctrl *gomock.Controller) *MockIntakeFormsClient {
	mock := &MockIntakeFormsClient{ctrl: ctrl}
	mock.recorder = &MockIntakeFormsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIntakeFormsClient) EXPECT() *MockIntakeFormsClientMockRecorder {
	return m.recorder
}

// AddIntakeForms mocks base method
func (m *MockIntakeFormsClient) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm, opts ...grpc.CallOption) (*intake_form.IntakeFormIdentifier, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddIntakeForms", varargs...)
	ret0, _ := ret[0].(*intake_form.IntakeFormIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIntakeForms indicates an expected call of AddIntakeForms
func (mr *MockIntakeFormsClientMockRecorder) AddIntakeForms(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIntakeForms", reflect.TypeOf((*MockIntakeFormsClient)(nil).AddIntakeForms), varargs...)
}

// UpdateIntakeForm mocks base method
func (m *MockIntakeFormsClient) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm, opts ...grpc.CallOption) (*empty.Empty, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateIntakeForm", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIntakeForm indicates an expected call of UpdateIntakeForm
func (mr *MockIntakeFormsClientMockRecorder) UpdateIntakeForm(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIntakeForm", reflect.TypeOf((*MockIntakeFormsClient)(nil).UpdateIntakeForm), varargs...)
}

// DeleteIntakeForm mocks base method
func (m *MockIntakeFormsClient) DeleteIntakeForm(ctx context.Context, in *intake_form.IntakeFormIdentifier, opts ...grpc.CallOption) (*empty.Empty, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteIntakeForm", varargs...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIntakeForm indicates an expected call of DeleteIntakeForm
func (mr *MockIntakeFormsClientMockRecorder) DeleteIntakeForm(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIntakeForm", reflect.TypeOf((*MockIntakeFormsClient)(nil).DeleteIntakeForm), varargs...)
}

// GetIntakeForms mocks base method
func (m *MockIntakeFormsClient) GetIntakeForms(ctx context.Context, in *locations.LocationRoot, opts ...grpc.CallOption) (*intake_form.IntakeFormList, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntakeForms", varargs...)
	ret0, _ := ret[0].(*intake_form.IntakeFormList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeForms indicates an expected call of GetIntakeForms
func (mr *MockIntakeFormsClientMockRecorder) GetIntakeForms(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeForms", reflect.TypeOf((*MockIntakeFormsClient)(nil).GetIntakeForms), varargs...)
}

// GetIntakeFormById mocks base method
func (m *MockIntakeFormsClient) GetIntakeFormById(ctx context.Context, in *intake_form.IntakeFormIdentifier, opts ...grpc.CallOption) (*intake_form.IntakeForm, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntakeFormById", varargs...)
	ret0, _ := ret[0].(*intake_form.IntakeForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeFormById indicates an expected call of GetIntakeFormById
func (mr *MockIntakeFormsClientMockRecorder) GetIntakeFormById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeFormById", reflect.TypeOf((*MockIntakeFormsClient)(nil).GetIntakeFormById), varargs...)
}

// MockIntakeFormsServer is a mock of IntakeFormsServer interface
type MockIntakeFormsServer struct {
	ctrl     *gomock.Controller
	recorder *MockIntakeFormsServerMockRecorder
}

// MockIntakeFormsServerMockRecorder is the mock recorder for MockIntakeFormsServer
type MockIntakeFormsServerMockRecorder struct {
	mock *MockIntakeFormsServer
}

// NewMockIntakeFormsServer creates a new mock instance
func NewMockIntakeFormsServer(ctrl *gomock.Controller) *MockIntakeFormsServer {
	mock := &MockIntakeFormsServer{ctrl: ctrl}
	mock.recorder = &MockIntakeFormsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIntakeFormsServer) EXPECT() *MockIntakeFormsServerMockRecorder {
	return m.recorder
}

// AddIntakeForms mocks base method
func (m *MockIntakeFormsServer) AddIntakeForms(arg0 context.Context, arg1 *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {
	ret := m.ctrl.Call(m, "AddIntakeForms", arg0, arg1)
	ret0, _ := ret[0].(*intake_form.IntakeFormIdentifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIntakeForms indicates an expected call of AddIntakeForms
func (mr *MockIntakeFormsServerMockRecorder) AddIntakeForms(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIntakeForms", reflect.TypeOf((*MockIntakeFormsServer)(nil).AddIntakeForms), arg0, arg1)
}

// UpdateIntakeForm mocks base method
func (m *MockIntakeFormsServer) UpdateIntakeForm(arg0 context.Context, arg1 *intake_form.IntakeForm) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "UpdateIntakeForm", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIntakeForm indicates an expected call of UpdateIntakeForm
func (mr *MockIntakeFormsServerMockRecorder) UpdateIntakeForm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIntakeForm", reflect.TypeOf((*MockIntakeFormsServer)(nil).UpdateIntakeForm), arg0, arg1)
}

// DeleteIntakeForm mocks base method
func (m *MockIntakeFormsServer) DeleteIntakeForm(arg0 context.Context, arg1 *intake_form.IntakeFormIdentifier) (*empty.Empty, error) {
	ret := m.ctrl.Call(m, "DeleteIntakeForm", arg0, arg1)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIntakeForm indicates an expected call of DeleteIntakeForm
func (mr *MockIntakeFormsServerMockRecorder) DeleteIntakeForm(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIntakeForm", reflect.TypeOf((*MockIntakeFormsServer)(nil).DeleteIntakeForm), arg0, arg1)
}

// GetIntakeForms mocks base method
func (m *MockIntakeFormsServer) GetIntakeForms(arg0 context.Context, arg1 *locations.LocationRoot) (*intake_form.IntakeFormList, error) {
	ret := m.ctrl.Call(m, "GetIntakeForms", arg0, arg1)
	ret0, _ := ret[0].(*intake_form.IntakeFormList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeForms indicates an expected call of GetIntakeForms
func (mr *MockIntakeFormsServerMockRecorder) GetIntakeForms(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeForms", reflect.TypeOf((*MockIntakeFormsServer)(nil).GetIntakeForms), arg0, arg1)
}

// GetIntakeFormById mocks base method
func (m *MockIntakeFormsServer) GetIntakeFormById(arg0 context.Context, arg1 *intake_form.IntakeFormIdentifier) (*intake_form.IntakeForm, error) {
	ret := m.ctrl.Call(m, "GetIntakeFormById", arg0, arg1)
	ret0, _ := ret[0].(*intake_form.IntakeForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntakeFormById indicates an expected call of GetIntakeFormById
func (mr *MockIntakeFormsServerMockRecorder) GetIntakeFormById(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntakeFormById", reflect.TypeOf((*MockIntakeFormsServer)(nil).GetIntakeFormById), arg0, arg1)
}