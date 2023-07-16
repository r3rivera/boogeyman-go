package hasher

import "golang.org/x/crypto/bcrypt"

type BCryptPassword struct {
	value string
}

func NewBCrypt(value string) BCryptPassword {
	return BCryptPassword{value}
}

func (b *BCryptPassword) HashItem() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(b.value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Provide the hash of the password and the plain password
func BCryptVerifyItem(hash1, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash1), []byte(password))
	if err == nil {
		return true
	}
	return false
}
