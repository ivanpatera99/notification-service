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

type MockNotificationRepository struct {
	mock.Mock
}

func (m *MockNotificationRepository) CheckIfNotificationIsSpam(notificationType string, userId string) (bool, error) {
	args := m.Called(notificationType, userId)
	return args.Bool(0), args.Error(1)
}

func (m *MockNotificationRepository) SaveNotificationEvent(notificationType string, userId string) error {
	args := m.Called(notificationType, userId)
	return args.Error(0)
}


func TestSendNotificationShouldReturnNil(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationRepo := new(MockNotificationRepository)
	notificationRepo.On("CheckIfNotificationIsSpam", "email", "123").Return(false, nil)
	notificationRepo.On("SaveNotificationEvent", "email", "123").Return(nil)
	notificationService.On("Send", "email", "123", "Hello, World!").Return(nil)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.Nil(t, err)
	notificationService.AssertExpectations(t)
	notificationRepo.AssertExpectations(t)
}

func TestSendNotificationThatIsSpamShouldReturnError(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationRepo := new(MockNotificationRepository)
	notificationRepo.On("CheckIfNotificationIsSpam", "email", "123").Return(true, nil)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.NotNil(t, err)
	assert.Equal(t, "NOTIFICATION_IS_SPAM", err.Error())
	notificationService.AssertExpectations(t)
	notificationRepo.AssertExpectations(t)
}

func TestSendNotificationFailsToCheckForSpamShouldReturnError(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationRepo := new(MockNotificationRepository)
	notificationRepo.On("CheckIfNotificationIsSpam", "email", "123").Return(false, assert.AnError)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.NotNil(t, err)
	assert.Equal(t, "ERROR_CHECKING_IF_NOTIFICATION_IS_SPAM", err.Error())
	notificationService.AssertExpectations(t)
	notificationRepo.AssertExpectations(t)
}

func TestSendNotificationFailsToSaveNotificationEventShouldReturnError(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationRepo := new(MockNotificationRepository)
	notificationRepo.On("CheckIfNotificationIsSpam", "email", "123").Return(false, nil)
	notificationRepo.On("SaveNotificationEvent", "email", "123").Return(assert.AnError)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.NotNil(t, err)
	assert.Equal(t, "ERROR_SAVING_NOTIFICATION_EVENT", err.Error())
	notificationService.AssertExpectations(t)
	notificationRepo.AssertExpectations(t)
}

func TestSendNotificationFailsToSendShouldReturnError(t *testing.T) {
	notificationService := new(MockNotificationService)
	notificationRepo := new(MockNotificationRepository)
	notificationRepo.On("CheckIfNotificationIsSpam", "email", "123").Return(false, nil)
	notificationRepo.On("SaveNotificationEvent", "email", "123").Return(nil)
	notificationService.On("Send", "email", "123", "Hello, World!").Return(assert.AnError)

	notificationsUseCase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	err := notificationsUseCase.SendNotification("email", "123", "Hello, World!")

	assert.NotNil(t, err)
	assert.Equal(t, "ERROR_SENDING_NOTIFICATION", err.Error())
	notificationService.AssertExpectations(t)
	notificationRepo.AssertExpectations(t)
}