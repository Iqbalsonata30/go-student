package domain

import "github.com/google/uuid"

type Student struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	IdentityNumber int       `json:"identityNumber"`
	Gender         string    `json:"gender"`
	Major          string    `json:"major"`
	Class          string    `json:"class"`
	Religion       string    `json:"religion"`
}
