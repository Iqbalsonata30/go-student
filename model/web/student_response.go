package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateStudentResponse struct {
	ID uuid.UUID `json:"id"`
}

type StudentResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	IdentityNumber int       `json:"identityNumber"`
	Gender         string    `json:"gender"`
	Major          string    `json:"major"`
	Class          string    `json:"class"`
	Religion       string    `json:"religion"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
