package player

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlayer(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	player := &Player{ID: "1", Email: "test@example.com", Name: "John"}

	mockRepo.On("GetByID", "1").Return(nil, nil)
	mockRepo.On("Create", player).Return(nil)

	err := service.CreatePlayer(player)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
