package service

import (
	"shockdart/internal/repository"
)

type NotificationRequest struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
	Target  string `json:"target"`
}

type NotificationService interface {
	SendNotification(req NotificationRequest) error
}

type notificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) SendNotification(req NotificationRequest) error {
	if req.Channel == "firebase" {
		return s.repo.EnqueueNotification(req)
	}
	return s.repo.EnqueueNotification(req)
}
