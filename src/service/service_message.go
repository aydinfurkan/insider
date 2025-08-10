package service

import (
	"context"
	"insider/src/db"
	"insider/src/domain"
	"time"
)

type MessageService struct {
	messageDb      *db.MessageDb
	webhookService *WebhookService
	messageCh      chan *domain.Message
	ctx            context.Context
	cancelCtx      context.CancelFunc
}

func NewMessageService(messageDb *db.MessageDb, webhookService *WebhookService) *MessageService {
	ctx, cancel := context.WithCancel(context.Background())
	s := &MessageService{
		messageDb:      messageDb,
		webhookService: webhookService,
		ctx:            ctx,
		cancelCtx:      cancel,
		messageCh:      make(chan *domain.Message, 2),
	}

	go s.listen()
	go s.start()
	go s.sendAll()

	return s
}

func (s *MessageService) Toggle() bool {
	if s.ctx.Err() == context.Canceled {
		s.ctx, s.cancelCtx = context.WithCancel(context.Background())
		go s.start()
		return true
	} else {
		s.cancelCtx()
		return false
	}
}

func (s *MessageService) start() error {
	ticker := time.NewTicker(time.Minute * 2)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return nil
		case <-ticker.C:
			messages, err := s.messageDb.GetPendingMessages(0, 2)
			if err != nil {
				return err
			}
			for _, message := range messages {
				s.messageCh <- &message
			}
			return nil
		}
	}
}

func (s *MessageService) sendAll() error {
	for {
		messages, err := s.messageDb.GetPendingMessages(0, 10)
		if err != nil {
			return err
		}
		if len(messages) == 0 {
			break
		}
		for _, message := range messages {
			s.messageCh <- &message
		}
	}
	return nil
}

func (s *MessageService) listen() {
	for message := range s.messageCh {
		_, err := s.webhookService.SendMessage(message.RecipientNumber, message.Content)
		if err != nil {
			continue
		}
		message.Sent()
		s.messageDb.UpdateMessage(message)
	}
}
