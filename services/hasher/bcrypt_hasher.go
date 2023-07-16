package hasher

import "golang.org/x/crypto/bcrypt"

type BCryptPassword struct {
	value string
}

func (b *BCryptPassword) HashItem() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(b.value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b *BCryptPassword) VerifyItem(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(b.value))
	return err == nil
}
