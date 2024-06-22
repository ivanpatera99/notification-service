package usecases

import (
	"errors"

	"github.com/ivanpatera99/notification-service/src/domain/ports"
)

type NotificationsUseCase struct {
	notificationService ports.NotificationService
	notificationRepo 	ports.NotificationRepository
}

func NewNotificationsUseCase(notificationService ports.NotificationService, notificationRepo ports.NotificationRepository) *NotificationsUseCase {
	return &NotificationsUseCase{notificationService: notificationService, notificationRepo: notificationRepo}
}

func (n *NotificationsUseCase) SendNotification(notificationType string, userId string, message string) error {
	isSpam, err := n.notificationRepo.CheckIfNotificationIsSpam(notificationType, userId)
	if err != nil {
		return errors.New("ERROR_CHECKING_IF_NOTIFICATION_IS_SPAM")
	}
	if isSpam {
		return errors.New("NOTIFICATION_IS_SPAM")
	}
	err = n.notificationRepo.SaveNotificationEvent(notificationType, userId)
	if err != nil {
		return errors.New("ERROR_SAVING_NOTIFICATION_EVENT")
	}
	err = n.notificationService.Send(notificationType, userId, message)
	if err != nil {
		return errors.New("ERROR_SENDING_NOTIFICATION")
	}
	return nil
}