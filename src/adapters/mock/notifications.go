package service_mocks

import (
	"errors"
	"math/rand"
)

type notificationServiceMock struct {}

func NewNotificationServiceMock() *notificationServiceMock {
	return &notificationServiceMock{}
}

func (n *notificationServiceMock) Send(notificationType string, userId string, message string) error {
	// Simulate a 1% chance of returning an error
	if rand.Float64() < 0.01 {
		return errors.New("mock_error_sending_notification")
	}
	// Send the notification successfully
	return nil
}