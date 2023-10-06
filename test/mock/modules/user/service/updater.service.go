// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/user/service/updater.service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"

	entity "gin-starter/entity"
)

// MockUserUpdaterUseCase is a mock of UserUpdaterUseCase interface.
type MockUserUpdaterUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUpdaterUseCaseMockRecorder
}

// MockUserUpdaterUseCaseMockRecorder is the mock recorder for MockUserUpdaterUseCase.
type MockUserUpdaterUseCaseMockRecorder struct {
	mock *MockUserUpdaterUseCase
}

// NewMockUserUpdaterUseCase creates a new mock instance.
func NewMockUserUpdaterUseCase(ctrl *gomock.Controller) *MockUserUpdaterUseCase {
	mock := &MockUserUpdaterUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUpdaterUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUpdaterUseCase) EXPECT() *MockUserUpdaterUseCaseMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserUpdaterUseCase) ChangePassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, userID, oldPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserUpdaterUseCaseMockRecorder) ChangePassword(ctx, userID, oldPassword, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).ChangePassword), ctx, userID, oldPassword, newPassword)
}

// ForgotPassword mocks base method.
func (m *MockUserUpdaterUseCase) ForgotPassword(ctx context.Context, userID uuid.UUID, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", ctx, userID, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockUserUpdaterUseCaseMockRecorder) ForgotPassword(ctx, userID, newPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).ForgotPassword), ctx, userID, newPassword)
}

// ForgotPasswordRequest mocks base method.
func (m *MockUserUpdaterUseCase) ForgotPasswordRequest(ctx context.Context, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPasswordRequest", ctx, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPasswordRequest indicates an expected call of ForgotPasswordRequest.
func (mr *MockUserUpdaterUseCaseMockRecorder) ForgotPasswordRequest(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPasswordRequest", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).ForgotPasswordRequest), ctx, email)
}

// ResendOTP mocks base method.
func (m *MockUserUpdaterUseCase) ResendOTP(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResendOTP", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResendOTP indicates an expected call of ResendOTP.
func (mr *MockUserUpdaterUseCaseMockRecorder) ResendOTP(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResendOTP", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).ResendOTP), ctx, userID)
}

// Update mocks base method.
func (m *MockUserUpdaterUseCase) Update(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserUpdaterUseCaseMockRecorder) Update(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).Update), ctx, user)
}

// VerifyOTP mocks base method.
func (m *MockUserUpdaterUseCase) VerifyOTP(ctx context.Context, userID uuid.UUID, otp string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyOTP", ctx, userID, otp)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyOTP indicates an expected call of VerifyOTP.
func (mr *MockUserUpdaterUseCaseMockRecorder) VerifyOTP(ctx, userID, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyOTP", reflect.TypeOf((*MockUserUpdaterUseCase)(nil).VerifyOTP), ctx, userID, otp)
}
