package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusPending string = "pending"
	StatusSent    string = "sent"
	StatusFailed  string = "failed"
)

type Message struct {
	Id              uint      `json:"id" gorm:"primaryKey; autoIncrement"`
	MessageId       uuid.UUID `json:"message_id" gorm:"type:uuid; default:uuid_generate_v4(); index:idx_message_id,unique"`
	Content         string    `json:"content" gorm:"not null"`
	RecipientNumber string    `json:"recipient_number" gorm:"not null"`
	Status          string    `json:"status" gorm:"not null; default:'pending'; index:idx_status"`
	ErrorCount      int       `json:"error_count" gorm:"default:0"`
	ErrorMessage    *string   `json:"error_message" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewMessage(content, recipientNumber string) *Message {
	return &Message{
		Content:         content,
		RecipientNumber: recipientNumber,
		Status:          string(StatusPending),
	}
}

func (m *Message) Sent() {
	m.Status = string(StatusSent)
}
