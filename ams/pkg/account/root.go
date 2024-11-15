package account

import (
	"time"

	"github.com/google/uuid"
)

type RootAccount struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Services     []Service
	CreationDate time.Time
}
