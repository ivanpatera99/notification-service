package ports

type NotificationService interface {
	Send(notificationType string, userId string, message string) error
}