package main

import (
	"github.com/ivanpatera99/notification-service/src/adapters/mock"
	"github.com/ivanpatera99/notification-service/src/adapters/sqlite"
	"github.com/ivanpatera99/notification-service/src/app"
	"github.com/ivanpatera99/notification-service/src/domain/usecases"
)

func main() {
	notificationRepo, err := sqlite.NewNotificationRepository()
	if err != nil {
		panic(err)
	}
	notificationService := service_mocks.NewNotificationServiceMock()
	notificationUsecase := usecases.NewNotificationsUseCase(notificationService, notificationRepo)
	app := app.NewApp(*notificationUsecase)
	app.Run()
}
