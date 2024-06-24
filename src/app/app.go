package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanpatera99/notification-service/src/domain/usecases"
)

// App struct
type App struct {
	notificationUseCase usecases.NotificationsUseCase
}

func NewApp(notificationUseCase usecases.NotificationsUseCase) *App {
	return &App{notificationUseCase: notificationUseCase}
}

func (a *App) Run() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/send-notification", func(c *gin.Context) {
		notificationType := c.PostForm("notification_type")
		userId := c.PostForm("user_id")
		message := c.PostForm("message")

		err := a.notificationUseCase.SendNotification(notificationType, userId, message)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Notification sent successfully",
		})
	})

	r.Run()
}