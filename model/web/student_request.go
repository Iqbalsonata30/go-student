package web

type StudentRequest struct {
	Name           string `validate:"required,min=1,max=100" json:"name"`
	IdentityNumber int    `validate:"required,numeric" json:"identityNumber"`
	Gender         string `validate:"required,min=1,max=20" json:"gender"`
	Major          string `validate:"required,min=1,max=50" json:"major"`
	Class          string `validate:"required,min=1,max=10" json:"class"`
	Religion       string `validate:"required,min=1,max=15" json:"religion"`
}
