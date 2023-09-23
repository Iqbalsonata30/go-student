package domain

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	ID             uuid.UUID
	Name           string
	IdentityNumber int
	Gender         string
	Major          string
	Class          string
	Religion       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
