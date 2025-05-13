package Models

import "time"

type LogEntry struct {
	ID        uint      `gorm:"primaryKey"`
	Timestamp time.Time `json:"timestamp" validate:"required"`
	Level     string    `json:"level" validate:"required,oneof=info warn error"`
	Source    string    `json:"source" validate:"required"`
	Message   string    `json:"message" validate:"required"`
	UserID    int       `json:"user_id,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	CreatedAt time.Time
}
