package message

import (
	"context"
	"time"
)

type service struct {
	Repository IRepository
	timeout    time.Duration
}

func NewService(messageRepo IRepository) IService {
	return &service{
		Repository: messageRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *service) GetMessageByEmail(ctx context.Context, username string) ([]Message, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) SaveMessage(message *Message) (*Message, error) {
	return nil, nil
}
