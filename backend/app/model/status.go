package model

import "time"

type Status struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	MessageID int       `gorm:"column:message_id;type:integer;not null" json:"message_id"`
	Status    string    `gorm:"column:status;type:varchar(20);not null" json:"status"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Status) TableName() string {
	return "statuses"
}
