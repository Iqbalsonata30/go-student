package web

type StudentRequest struct {
	Name           string `json:"name"`
	IdentityNumber int    `json:"identityNumber"`
	Gender         string `json:"gender"`
	Major          string `json:"major"`
	Class          string `json:"class"`
	Religion       string `json:"religion"`
}
