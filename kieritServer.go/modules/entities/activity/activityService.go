package activity

import "time"

type ActivityService struct {
	Repository *ActivityRepo
	timeout    time.Duration
}

func NewMessageService(messageRepo *ActivityRepo) *ActivityService {
	return &ActivityService{
		Repository: messageRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}
