package repository

import (
	"awesomeProject7/models"
	"github.com/jmoiron/sqlx"
)

type Subscription interface {
	Create(email string) error
	GetByEmail(email string) error
	GetAll(email string) ([]models.Subscription, error)
}

type Repository struct {
	Subscription
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Subscription: NewSubscriptionPostgres(db)}
}
