package db

import (
	"insider/src/config"
	"insider/src/domain"

	"gorm.io/gorm"
)

type MessageDb struct {
	db *gorm.DB
}

func NewMessageDb(cfg *config.ConfigType) *MessageDb {
	db, err := connectPostgre(cfg.POSTGREDB_URL, &domain.Message{})
	if err != nil {
		panic(err)
	}
	return &MessageDb{db: db}
}

func (m *MessageDb) CreateMessage(message *domain.Message) error {
	return m.db.Create(message).Error
}

func (m *MessageDb) GetSentMessages(offset, limit int) ([]domain.Message, error) {
	var messages []domain.Message
	err := m.db.Where("status = ?", "sent").Offset(offset).Limit(limit).Find(&messages).Error
	return messages, err
}

func (m *MessageDb) GetPendingMessages(offset, limit int) ([]domain.Message, error) {
	var messages []domain.Message
	err := m.db.Where("status = ?", "pending").Order("created_at asc").Offset(offset).Limit(limit).Find(&messages).Error
	return messages, err
}

func (m *MessageDb) UpdateMessage(message *domain.Message) error {
	return m.db.Save(message).Error
}
