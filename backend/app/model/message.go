package model

import "time"

type Message struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	SenderID   int       `gorm:"column:sender_id;type:integer;not null" json:"sender_id"`
	ReceiverID int       `gorm:"column:receiver_id;type:integer;not null" json:"receiver_id"`
	Content    string    `gorm:"column:content;type:text;not null" json:"content"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}


func (Message) TableName() string {
	return "messages"
}