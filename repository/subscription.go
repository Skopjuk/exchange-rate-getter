package repository

import (
	"awesomeProject7/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type SubscriptionPostgres struct {
	db *sqlx.DB
}

func NewSubscriptionPostgres(db *sqlx.DB) *SubscriptionPostgres {
	return &SubscriptionPostgres{db: db}
}

func (s *SubscriptionPostgres) Create(email string) error {
	query := fmt.Sprintf("INSERT INTO %s (email) VALUES ($1)", subscriptionsTable)
	_, err := s.db.Query(query, email)
	if err != nil {
		logrus.Errorf("error while inserting to db: %s", err)
		return err
	}

	return nil
}

func (s *SubscriptionPostgres) GetByEmail(email string) error {
	var response []models.Subscription
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1 LIMIT 1", subscriptionsTable)
	err := s.db.Get(&response, query, email)
	if err != nil {
		return errors.New("email is not found")
	}

	return nil
}

func (s *SubscriptionPostgres) GetAll(email string) ([]models.Subscription, error) {
	var response []models.Subscription
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1", subscriptionsTable)
	err := s.db.Get(&response, query, email)
	if err != nil {
		return response, errors.New("email is not found")
	}

	return response, nil
}
