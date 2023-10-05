package web

import "github.com/google/uuid"

type CreateUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UserLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
