package account

import (
	"time"

	"github.com/google/uuid"
)

type ServiceAccount struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Groups       []ServiceAccountGroup
	CreationDate time.Time
}

func NewServiceAccount(username string, email string) *ServiceAccount {
	return &ServiceAccount{
		ID:       uuid.New(),
		Username: username,
		Email:    email,
	}
}
