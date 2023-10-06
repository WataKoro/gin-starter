// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/auth/repository/auth.repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	entity "gin-starter/entity"
)

// MockAuthRepositoryUseCase is a mock of AuthRepositoryUseCase interface.
type MockAuthRepositoryUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryUseCaseMockRecorder
}

// MockAuthRepositoryUseCaseMockRecorder is the mock recorder for MockAuthRepositoryUseCase.
type MockAuthRepositoryUseCaseMockRecorder struct {
	mock *MockAuthRepositoryUseCase
}

// NewMockAuthRepositoryUseCase creates a new mock instance.
func NewMockAuthRepositoryUseCase(ctrl *gomock.Controller) *MockAuthRepositoryUseCase {
	mock := &MockAuthRepositoryUseCase{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepositoryUseCase) EXPECT() *MockAuthRepositoryUseCaseMockRecorder {
	return m.recorder
}

// GetAdminByEmail mocks base method.
func (m *MockAuthRepositoryUseCase) GetAdminByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminByEmail indicates an expected call of GetAdminByEmail.
func (mr *MockAuthRepositoryUseCaseMockRecorder) GetAdminByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminByEmail", reflect.TypeOf((*MockAuthRepositoryUseCase)(nil).GetAdminByEmail), ctx, email)
}

// GetUserByEmail mocks base method.
func (m *MockAuthRepositoryUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockAuthRepositoryUseCaseMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockAuthRepositoryUseCase)(nil).GetUserByEmail), ctx, email)
}

// UpdateOTP mocks base method.
func (m *MockAuthRepositoryUseCase) UpdateOTP(ctx context.Context, user *entity.User, otp string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOTP", ctx, user, otp)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOTP indicates an expected call of UpdateOTP.
func (mr *MockAuthRepositoryUseCaseMockRecorder) UpdateOTP(ctx, user, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOTP", reflect.TypeOf((*MockAuthRepositoryUseCase)(nil).UpdateOTP), ctx, user, otp)
}
