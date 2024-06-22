package usecases

import (
	"github.com/ivanpatera99/notification-service/src/domain/ports"
)

type NotificationsUseCase struct {
	notificationService ports.NotificationService
}

func NewNotificationsUseCase(notificationService ports.NotificationService) *NotificationsUseCase {
	return &NotificationsUseCase{notificationService: notificationService}
}

func (n *NotificationsUseCase) SendNotification(notificationType string, userId string, message string) error {
	return n.notificationService.Send(notificationType, userId, message)
}