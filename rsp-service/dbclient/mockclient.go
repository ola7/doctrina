package dbclient

import (
	"../model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient is a mock db client that does the bare minimum
type MockBoltClient struct {
	mock.Mock
}

// QueryUser returns dummy user
func (m *MockBoltClient) QueryUser(userID string) (model.User, error) {
	args := m.Mock.Called(userID)
	return args.Get(0).(model.User), args.Error(1)
}

// OpenBoltDb does nothing
func (m *MockBoltClient) OpenBoltDb() {
}

// SeedFakeUsers does nothing
func (m *MockBoltClient) SeedFakeUsers(n int) {
}

// CheckStatus returns if client has already been called
func (m *MockBoltClient) CheckStatus() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}
