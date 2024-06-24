package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	driverName = "sqlite3"
	dbPath     = "db.db"
)

type NotificationRepository struct {
	conn *sql.DB
}

func NewNotificationRepository() (*NotificationRepository, error) {
	conn, err := sql.Open(driverName, dbPath)
	if err != nil {
		return nil, err
	}
	return &NotificationRepository{conn: conn}, nil
}

func (n *NotificationRepository) CheckIfNotificationIsSpam(notificationType string, userId string) (bool, error) {
	// Check if the notification is spam
	// First version will check if user has been notfified two times in the last 10 seconds
	rows, err := n.conn.Query("SELECT COUNT(*) FROM notifications WHERE user_id = ? AND notification_type = ?", userId, notificationType)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return false, err
		}
	}
	if count >= 2 {
		return true, nil
	}
	return false, nil
}

func (n *NotificationRepository) SaveNotificationEvent(notificationType string, userId string) error {
	// Save the notification event
	_, err := n.conn.Exec("INSERT INTO notifications (notification_type, user_id) VALUES (?, ?)", notificationType, userId)
	if err != nil {
		return err
	}
	return nil
}
