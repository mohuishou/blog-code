// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mohuishou/new-project/internal/domain (interfaces: IArticleRepo)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/mohuishou/new-project/internal/domain"
)

// MockIArticleRepo is a mock of IArticleRepo interface.
type MockIArticleRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleRepoMockRecorder
}

// MockIArticleRepoMockRecorder is the mock recorder for MockIArticleRepo.
type MockIArticleRepoMockRecorder struct {
	mock *MockIArticleRepo
}

// NewMockIArticleRepo creates a new mock instance.
func NewMockIArticleRepo(ctrl *gomock.Controller) *MockIArticleRepo {
	mock := &MockIArticleRepo{ctrl: ctrl}
	mock.recorder = &MockIArticleRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleRepo) EXPECT() *MockIArticleRepoMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockIArticleRepo) CreateArticle(arg0 context.Context, arg1 *domain.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockIArticleRepoMockRecorder) CreateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIArticleRepo)(nil).CreateArticle), arg0, arg1)
}

// CreateArticleTags mocks base method.
func (m *MockIArticleRepo) CreateArticleTags(arg0 context.Context, arg1 []*domain.ArticleTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticleTags", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateArticleTags indicates an expected call of CreateArticleTags.
func (mr *MockIArticleRepoMockRecorder) CreateArticleTags(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticleTags", reflect.TypeOf((*MockIArticleRepo)(nil).CreateArticleTags), arg0, arg1)
}

// GetArticle mocks base method.
func (m *MockIArticleRepo) GetArticle(arg0 context.Context, arg1 ...domain.DBOption) (*domain.Article, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetArticle", varargs...)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockIArticleRepoMockRecorder) GetArticle(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockIArticleRepo)(nil).GetArticle), varargs...)
}

// GetArticleByID mocks base method.
func (m *MockIArticleRepo) GetArticleByID(arg0 context.Context, arg1 int) (*domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", arg0, arg1)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockIArticleRepoMockRecorder) GetArticleByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockIArticleRepo)(nil).GetArticleByID), arg0, arg1)
}

// GetArticleByTitle mocks base method.
func (m *MockIArticleRepo) GetArticleByTitle(arg0 context.Context, arg1 string) (*domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByTitle", arg0, arg1)
	ret0, _ := ret[0].(*domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByTitle indicates an expected call of GetArticleByTitle.
func (mr *MockIArticleRepoMockRecorder) GetArticleByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByTitle", reflect.TypeOf((*MockIArticleRepo)(nil).GetArticleByTitle), arg0, arg1)
}

// Tx mocks base method.
func (m *MockIArticleRepo) Tx(arg0 context.Context, arg1 func(context.Context, domain.IArticleRepo) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tx indicates an expected call of Tx.
func (mr *MockIArticleRepoMockRecorder) Tx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockIArticleRepo)(nil).Tx), arg0, arg1)
}

// WithByID mocks base method.
func (m *MockIArticleRepo) WithByID(arg0 uint) domain.DBOption {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithByID", arg0)
	ret0, _ := ret[0].(domain.DBOption)
	return ret0
}

// WithByID indicates an expected call of WithByID.
func (mr *MockIArticleRepoMockRecorder) WithByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithByID", reflect.TypeOf((*MockIArticleRepo)(nil).WithByID), arg0)
}

// WithByTitle mocks base method.
func (m *MockIArticleRepo) WithByTitle(arg0 string) domain.DBOption {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithByTitle", arg0)
	ret0, _ := ret[0].(domain.DBOption)
	return ret0
}

// WithByTitle indicates an expected call of WithByTitle.
func (mr *MockIArticleRepoMockRecorder) WithByTitle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithByTitle", reflect.TypeOf((*MockIArticleRepo)(nil).WithByTitle), arg0)
}