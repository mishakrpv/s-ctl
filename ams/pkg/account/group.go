package account

import (
	"time"

	"github.com/google/uuid"
)

type ServiceAccountGroup struct {
	ID           uuid.UUID
	Name         string
	Accounts     []ServiceAccount
	Policy       *GroupPolicy
	CreationDate time.Time
}

type GroupPolicy struct {
	ID uuid.UUID
}
