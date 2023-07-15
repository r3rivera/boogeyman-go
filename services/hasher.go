package services

import "golang.org/x/crypto/bcrypt"

type BCryptPassword struct {
	value string
}

type Encryptor interface {
	HashPassword() (string, error)
	VerifyPassword(hash string) bool
}

func (b *BCryptPassword) HashPassword() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(b.value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b *BCryptPassword) VerifyPassword(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(b.value))
	return err == nil
}
