// Code generated by MockGen. DO NOT EDIT.
// Source: internal/entity/manager/audit.go

// Package manager is a generated GoMock package.
package manager

import (
	reflect "reflect"

	model "github.com/eko/authz/backend/internal/entity/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAudit is a mock of Audit interface.
type MockAudit struct {
	ctrl     *gomock.Controller
	recorder *MockAuditMockRecorder
}

// MockAuditMockRecorder is the mock recorder for MockAudit.
type MockAuditMockRecorder struct {
	mock *MockAudit
}

// NewMockAudit creates a new mock instance.
func NewMockAudit(ctrl *gomock.Controller) *MockAudit {
	mock := &MockAudit{ctrl: ctrl}
	mock.recorder = &MockAuditMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAudit) EXPECT() *MockAuditMockRecorder {
	return m.recorder
}

// BatchAdd mocks base method.
func (m *MockAudit) BatchAdd(audits []*model.Audit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchAdd", audits)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchAdd indicates an expected call of BatchAdd.
func (mr *MockAuditMockRecorder) BatchAdd(audits interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchAdd", reflect.TypeOf((*MockAudit)(nil).BatchAdd), audits)
}

// GetRepository mocks base method.
func (m *MockAudit) GetRepository() AuditRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRepository")
	ret0, _ := ret[0].(AuditRepository)
	return ret0
}

// GetRepository indicates an expected call of GetRepository.
func (mr *MockAuditMockRecorder) GetRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepository", reflect.TypeOf((*MockAudit)(nil).GetRepository))
}