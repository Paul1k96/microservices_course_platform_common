// Code generated by MockGen. DO NOT EDIT.
// Source: db.go
//
// Generated by this command:
//
//	mockgen -source db.go -destination mocks/db.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	db "github.com/Paul1k96/microservices_course_platform_common/pkg/client/db"
	pgconn "github.com/jackc/pgconn"
	v4 "github.com/jackc/pgx/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// DB mocks base method.
func (m *MockClient) DB() db.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(db.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockClientMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockClient)(nil).DB))
}

// MockTxManager is a mock of TxManager interface.
type MockTxManager struct {
	ctrl     *gomock.Controller
	recorder *MockTxManagerMockRecorder
	isgomock struct{}
}

// MockTxManagerMockRecorder is the mock recorder for MockTxManager.
type MockTxManagerMockRecorder struct {
	mock *MockTxManager
}

// NewMockTxManager creates a new mock instance.
func NewMockTxManager(ctrl *gomock.Controller) *MockTxManager {
	mock := &MockTxManager{ctrl: ctrl}
	mock.recorder = &MockTxManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxManager) EXPECT() *MockTxManagerMockRecorder {
	return m.recorder
}

// ReadCommitted mocks base method.
func (m *MockTxManager) ReadCommitted(ctx context.Context, f db.Handler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCommitted", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadCommitted indicates an expected call of ReadCommitted.
func (mr *MockTxManagerMockRecorder) ReadCommitted(ctx, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCommitted", reflect.TypeOf((*MockTxManager)(nil).ReadCommitted), ctx, f)
}

// MockTransactor is a mock of Transactor interface.
type MockTransactor struct {
	ctrl     *gomock.Controller
	recorder *MockTransactorMockRecorder
	isgomock struct{}
}

// MockTransactorMockRecorder is the mock recorder for MockTransactor.
type MockTransactorMockRecorder struct {
	mock *MockTransactor
}

// NewMockTransactor creates a new mock instance.
func NewMockTransactor(ctrl *gomock.Controller) *MockTransactor {
	mock := &MockTransactor{ctrl: ctrl}
	mock.recorder = &MockTransactorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactor) EXPECT() *MockTransactorMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockTransactor) BeginTx(ctx context.Context, txOptions v4.TxOptions) (v4.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(v4.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockTransactorMockRecorder) BeginTx(ctx, txOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockTransactor)(nil).BeginTx), ctx, txOptions)
}

// MockSQLExecutor is a mock of SQLExecutor interface.
type MockSQLExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockSQLExecutorMockRecorder
	isgomock struct{}
}

// MockSQLExecutorMockRecorder is the mock recorder for MockSQLExecutor.
type MockSQLExecutorMockRecorder struct {
	mock *MockSQLExecutor
}

// NewMockSQLExecutor creates a new mock instance.
func NewMockSQLExecutor(ctrl *gomock.Controller) *MockSQLExecutor {
	mock := &MockSQLExecutor{ctrl: ctrl}
	mock.recorder = &MockSQLExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLExecutor) EXPECT() *MockSQLExecutorMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockSQLExecutor) ExecContext(ctx context.Context, q db.Query, args ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockSQLExecutorMockRecorder) ExecContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockSQLExecutor)(nil).ExecContext), varargs...)
}

// QueryContext mocks base method.
func (m *MockSQLExecutor) QueryContext(ctx context.Context, q db.Query, args ...any) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockSQLExecutorMockRecorder) QueryContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockSQLExecutor)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockSQLExecutor) QueryRowContext(ctx context.Context, q db.Query, args ...any) v4.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockSQLExecutorMockRecorder) QueryRowContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockSQLExecutor)(nil).QueryRowContext), varargs...)
}

// ScanAllContext mocks base method.
func (m *MockSQLExecutor) ScanAllContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAllContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAllContext indicates an expected call of ScanAllContext.
func (mr *MockSQLExecutorMockRecorder) ScanAllContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAllContext", reflect.TypeOf((*MockSQLExecutor)(nil).ScanAllContext), varargs...)
}

// ScanOneContext mocks base method.
func (m *MockSQLExecutor) ScanOneContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOneContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOneContext indicates an expected call of ScanOneContext.
func (mr *MockSQLExecutorMockRecorder) ScanOneContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOneContext", reflect.TypeOf((*MockSQLExecutor)(nil).ScanOneContext), varargs...)
}

// MockNamedExecutor is a mock of NamedExecutor interface.
type MockNamedExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockNamedExecutorMockRecorder
	isgomock struct{}
}

// MockNamedExecutorMockRecorder is the mock recorder for MockNamedExecutor.
type MockNamedExecutorMockRecorder struct {
	mock *MockNamedExecutor
}

// NewMockNamedExecutor creates a new mock instance.
func NewMockNamedExecutor(ctrl *gomock.Controller) *MockNamedExecutor {
	mock := &MockNamedExecutor{ctrl: ctrl}
	mock.recorder = &MockNamedExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamedExecutor) EXPECT() *MockNamedExecutorMockRecorder {
	return m.recorder
}

// ScanAllContext mocks base method.
func (m *MockNamedExecutor) ScanAllContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAllContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAllContext indicates an expected call of ScanAllContext.
func (mr *MockNamedExecutorMockRecorder) ScanAllContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAllContext", reflect.TypeOf((*MockNamedExecutor)(nil).ScanAllContext), varargs...)
}

// ScanOneContext mocks base method.
func (m *MockNamedExecutor) ScanOneContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOneContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOneContext indicates an expected call of ScanOneContext.
func (mr *MockNamedExecutorMockRecorder) ScanOneContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOneContext", reflect.TypeOf((*MockNamedExecutor)(nil).ScanOneContext), varargs...)
}

// MockQueryExecutor is a mock of QueryExecutor interface.
type MockQueryExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockQueryExecutorMockRecorder
	isgomock struct{}
}

// MockQueryExecutorMockRecorder is the mock recorder for MockQueryExecutor.
type MockQueryExecutorMockRecorder struct {
	mock *MockQueryExecutor
}

// NewMockQueryExecutor creates a new mock instance.
func NewMockQueryExecutor(ctrl *gomock.Controller) *MockQueryExecutor {
	mock := &MockQueryExecutor{ctrl: ctrl}
	mock.recorder = &MockQueryExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryExecutor) EXPECT() *MockQueryExecutorMockRecorder {
	return m.recorder
}

// ExecContext mocks base method.
func (m *MockQueryExecutor) ExecContext(ctx context.Context, q db.Query, args ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockQueryExecutorMockRecorder) ExecContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockQueryExecutor)(nil).ExecContext), varargs...)
}

// QueryContext mocks base method.
func (m *MockQueryExecutor) QueryContext(ctx context.Context, q db.Query, args ...any) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockQueryExecutorMockRecorder) QueryContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockQueryExecutor)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockQueryExecutor) QueryRowContext(ctx context.Context, q db.Query, args ...any) v4.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockQueryExecutorMockRecorder) QueryRowContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockQueryExecutor)(nil).QueryRowContext), varargs...)
}

// MockPinger is a mock of Pinger interface.
type MockPinger struct {
	ctrl     *gomock.Controller
	recorder *MockPingerMockRecorder
	isgomock struct{}
}

// MockPingerMockRecorder is the mock recorder for MockPinger.
type MockPingerMockRecorder struct {
	mock *MockPinger
}

// NewMockPinger creates a new mock instance.
func NewMockPinger(ctrl *gomock.Controller) *MockPinger {
	mock := &MockPinger{ctrl: ctrl}
	mock.recorder = &MockPingerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPinger) EXPECT() *MockPingerMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockPinger) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPingerMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPinger)(nil).Ping), ctx)
}

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
	isgomock struct{}
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockDB) BeginTx(ctx context.Context, txOptions v4.TxOptions) (v4.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(v4.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDBMockRecorder) BeginTx(ctx, txOptions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDB)(nil).BeginTx), ctx, txOptions)
}

// Close mocks base method.
func (m *MockDB) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}

// ExecContext mocks base method.
func (m *MockDB) ExecContext(ctx context.Context, q db.Query, args ...any) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockDBMockRecorder) ExecContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockDB)(nil).ExecContext), varargs...)
}

// Ping mocks base method.
func (m *MockDB) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDBMockRecorder) Ping(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDB)(nil).Ping), ctx)
}

// QueryContext mocks base method.
func (m *MockDB) QueryContext(ctx context.Context, q db.Query, args ...any) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryContext", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryContext indicates an expected call of QueryContext.
func (mr *MockDBMockRecorder) QueryContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryContext", reflect.TypeOf((*MockDB)(nil).QueryContext), varargs...)
}

// QueryRowContext mocks base method.
func (m *MockDB) QueryRowContext(ctx context.Context, q db.Query, args ...any) v4.Row {
	m.ctrl.T.Helper()
	varargs := []any{ctx, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRowContext", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRowContext indicates an expected call of QueryRowContext.
func (mr *MockDBMockRecorder) QueryRowContext(ctx, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRowContext", reflect.TypeOf((*MockDB)(nil).QueryRowContext), varargs...)
}

// ScanAllContext mocks base method.
func (m *MockDB) ScanAllContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanAllContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanAllContext indicates an expected call of ScanAllContext.
func (mr *MockDBMockRecorder) ScanAllContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanAllContext", reflect.TypeOf((*MockDB)(nil).ScanAllContext), varargs...)
}

// ScanOneContext mocks base method.
func (m *MockDB) ScanOneContext(ctx context.Context, dest any, q db.Query, args ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, dest, q}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScanOneContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScanOneContext indicates an expected call of ScanOneContext.
func (mr *MockDBMockRecorder) ScanOneContext(ctx, dest, q any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, dest, q}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanOneContext", reflect.TypeOf((*MockDB)(nil).ScanOneContext), varargs...)
}
