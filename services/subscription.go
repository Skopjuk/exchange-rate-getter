package services

import (
	"awesomeProject7/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) CreateSubscription(email string) error {
	return s.repo.Create(email)
}

func (s *SubscriptionService) GetSubscriptions(email string) error {
	return s.repo.GetByEmail(email)
}
