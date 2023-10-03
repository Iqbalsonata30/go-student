package web

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
