package web

import (
	"github.com/google/uuid"
)

type StudentResponse struct {
	ID uuid.UUID `json:"id"`
}
