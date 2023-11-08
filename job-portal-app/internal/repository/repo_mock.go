// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go
//
// Generated by this command:
//
//	mockgen -source=repo.go -destination=repo_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	models "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// FetchAllCompanies mocks base method.
func (m *MockUserRepo) FetchAllCompanies(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllCompanies", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllCompanies indicates an expected call of FetchAllCompanies.
func (mr *MockUserRepoMockRecorder) FetchAllCompanies(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllCompanies", reflect.TypeOf((*MockUserRepo)(nil).FetchAllCompanies), ctx)
}

// FetchAllJobPostings mocks base method.
func (m *MockUserRepo) FetchAllJobPostings(ctx context.Context) ([]models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAllJobPostings", ctx)
	ret0, _ := ret[0].([]models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchAllJobPostings indicates an expected call of FetchAllJobPostings.
func (mr *MockUserRepoMockRecorder) FetchAllJobPostings(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAllJobPostings", reflect.TypeOf((*MockUserRepo)(nil).FetchAllJobPostings), ctx)
}

// FetchCompanyByID mocks base method.
func (m *MockUserRepo) FetchCompanyByID(ctx context.Context, cid uint64) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCompanyByID", ctx, cid)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCompanyByID indicates an expected call of FetchCompanyByID.
func (mr *MockUserRepoMockRecorder) FetchCompanyByID(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCompanyByID", reflect.TypeOf((*MockUserRepo)(nil).FetchCompanyByID), ctx, cid)
}

// FetchJobPostingByID mocks base method.
func (m *MockUserRepo) FetchJobPostingByID(ctx context.Context, jid uint64) (models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobPostingByID", ctx, jid)
	ret0, _ := ret[0].(models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobPostingByID indicates an expected call of FetchJobPostingByID.
func (mr *MockUserRepoMockRecorder) FetchJobPostingByID(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobPostingByID", reflect.TypeOf((*MockUserRepo)(nil).FetchJobPostingByID), ctx, jid)
}

// FetchJobsForCompany mocks base method.
func (m *MockUserRepo) FetchJobsForCompany(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobsForCompany", ctx, cid)
	ret0, _ := ret[0].([]models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobsForCompany indicates an expected call of FetchJobsForCompany.
func (mr *MockUserRepoMockRecorder) FetchJobsForCompany(ctx, cid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobsForCompany", reflect.TypeOf((*MockUserRepo)(nil).FetchJobsForCompany), ctx, cid)
}

// InsertCompany mocks base method.
func (m *MockUserRepo) InsertCompany(ctx context.Context, companyData models.Company) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCompany", ctx, companyData)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCompany indicates an expected call of InsertCompany.
func (mr *MockUserRepoMockRecorder) InsertCompany(ctx, companyData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCompany", reflect.TypeOf((*MockUserRepo)(nil).InsertCompany), ctx, companyData)
}

// InsertJobPosting mocks base method.
func (m *MockUserRepo) InsertJobPosting(ctx context.Context, jobData models.Jobs) (models.Jobs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertJobPosting", ctx, jobData)
	ret0, _ := ret[0].(models.Jobs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertJobPosting indicates an expected call of InsertJobPosting.
func (mr *MockUserRepoMockRecorder) InsertJobPosting(ctx, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertJobPosting", reflect.TypeOf((*MockUserRepo)(nil).InsertJobPosting), ctx, jobData)
}

// InsertUser mocks base method.
func (m *MockUserRepo) InsertUser(ctx context.Context, userData models.User) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, userData)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockUserRepoMockRecorder) InsertUser(ctx, userData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserRepo)(nil).InsertUser), ctx, userData)
}

// VerifyUserCredentials mocks base method.
func (m *MockUserRepo) VerifyUserCredentials(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserCredentials", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyUserCredentials indicates an expected call of VerifyUserCredentials.
func (mr *MockUserRepoMockRecorder) VerifyUserCredentials(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserCredentials", reflect.TypeOf((*MockUserRepo)(nil).VerifyUserCredentials), ctx, email)
}
