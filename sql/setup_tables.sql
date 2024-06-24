-- Create a new table called 'notifications'
CREATE TABLE IF NOT EXISTS notifications (
    user_id INTEGER,
    notification_type TEXT,
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notifications_configurations (
    notification_type TEXT,
    sends_per_day INTEGER,
)