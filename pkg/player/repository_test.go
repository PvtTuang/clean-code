package player

import "github.com/stretchr/testify/mock"

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(player *Player) error {
	args := m.Called(player)
	return args.Error(0)
}

func (m *MockRepository) GetByID(id string) (*Player, error) {
	args := m.Called(id)
	if obj, ok := args.Get(0).(*Player); ok {
		return obj, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockRepository) Update(player *Player) error {
	args := m.Called(player)
	return args.Error(0)
}

func (m *MockRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRepository) ListAll() ([]Player, error) {
	args := m.Called()
	if obj, ok := args.Get(0).([]Player); ok {
		return obj, args.Error(1)
	}
	return nil, args.Error(1)
}
