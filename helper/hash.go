package helper

import "golang.org/x/crypto/bcrypt"

func HashAndSalted(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, err
}
