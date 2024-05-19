package services

import (
	"awesomeProject7/repository"
)

type Subscription interface {
	CreateSubscription(email string) error
	GetSubscriptions(email string) error
}

type Rate interface {
	GetExchangeRate() (float32, error)
}

type Service struct {
	Subscription
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Subscription: NewSubscriptionService(repo)}
}
