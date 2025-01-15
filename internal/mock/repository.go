package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/ainurqa95/mood-lifter/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRepository) Create(ctx context.Context, userUUID string, info *model.UserInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userUUID, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserRepositoryMockRecorder) Create(ctx, userUUID, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRepository)(nil).Create), ctx, userUUID, info)
}

// GetByLimitOffset mocks base method.
func (m *MockUserRepository) GetByLimitOffset(ctx context.Context, limit, offset int) ([]model.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPeriodWithLimitOffset", ctx, limit, offset)
	ret0, _ := ret[0].([]model.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLimitOffset indicates an expected call of GetByLimitOffset.
func (mr *MockUserRepositoryMockRecorder) GetByLimitOffset(ctx, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPeriodWithLimitOffset", reflect.TypeOf((*MockUserRepository)(nil).GetByLimitOffset), ctx, limit, offset)
}

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageRepository) Create(ctx context.Context, message model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessageRepositoryMockRecorder) Create(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepository)(nil).Create), ctx, message)
}

// MockComplimentRepository is a mock of ComplimentRepository interface.
type MockComplimentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockComplimentRepositoryMockRecorder
}

// MockComplimentRepositoryMockRecorder is the mock recorder for MockComplimentRepository.
type MockComplimentRepositoryMockRecorder struct {
	mock *MockComplimentRepository
}

// NewMockComplimentRepository creates a new mock instance.
func NewMockComplimentRepository(ctrl *gomock.Controller) *MockComplimentRepository {
	mock := &MockComplimentRepository{ctrl: ctrl}
	mock.recorder = &MockComplimentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComplimentRepository) EXPECT() *MockComplimentRepositoryMockRecorder {
	return m.recorder
}

// GetRandom mocks base method.
func (m *MockComplimentRepository) GetRandom(ctx context.Context) (*model.Compliment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandom", ctx)
	ret0, _ := ret[0].(*model.Compliment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandom indicates an expected call of GetRandom.
func (mr *MockComplimentRepositoryMockRecorder) GetRandom(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandom", reflect.TypeOf((*MockComplimentRepository)(nil).GetRandom), ctx)
}
