package ports

type NotificationService interface {
	Send(notificationType string, userId string, message string) error
}

type NotificationRepository interface {
	CheckIfNotificationIsSpam(notificationType string, userId string) (bool, error)
	SaveNotificationEvent(notificationType string, userId string) error
}