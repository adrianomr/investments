// Code generated by MockGen. DO NOT EDIT.
// Source: investment_get_all.go
//
// Generated by this command:
//
//	mockgen -source investment_get_all.go -destination mock/investment_get_all_mock.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/adrianomr/investments/src/domain/models"
	gomock "go.uber.org/mock/gomock"
)

// MockIInvestmentGetAll is a mock of IInvestmentGetAll interface.
type MockIInvestmentGetAll struct {
	ctrl     *gomock.Controller
	recorder *MockIInvestmentGetAllMockRecorder
}

// MockIInvestmentGetAllMockRecorder is the mock recorder for MockIInvestmentGetAll.
type MockIInvestmentGetAllMockRecorder struct {
	mock *MockIInvestmentGetAll
}

// NewMockIInvestmentGetAll creates a new mock instance.
func NewMockIInvestmentGetAll(ctrl *gomock.Controller) *MockIInvestmentGetAll {
	mock := &MockIInvestmentGetAll{ctrl: ctrl}
	mock.recorder = &MockIInvestmentGetAllMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInvestmentGetAll) EXPECT() *MockIInvestmentGetAllMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIInvestmentGetAll) Execute(ctx context.Context) ([]models.Investment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx)
	ret0, _ := ret[0].([]models.Investment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIInvestmentGetAllMockRecorder) Execute(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIInvestmentGetAll)(nil).Execute), ctx)
}
