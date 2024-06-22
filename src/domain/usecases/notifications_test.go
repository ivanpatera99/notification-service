package usecases_test

import (
	"testing"

	"github.com/ivanpatera99/notification-service/src/domain/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNotificationService struct {
	mock.Mock
}

func (m *MockNotificationService) Send(notificationType string, userId string, message string) error {
	args := m.Called(notificationType, userId, message)
	return args.Error(0)
}

func TestSendNotification(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationService.On("Send", "email", "123", "Hello, World!").Return(nil)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.Nil(t, err)
	notificationService.AssertExpectations(t)
}
