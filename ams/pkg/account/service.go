package account

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID              uuid.UUID
	RootID          uuid.UUID
	Name            string
	ServiceAccounts []ServiceAccount
	Groups          []ServiceAccountGroup
	CreationDate    time.Time
}
