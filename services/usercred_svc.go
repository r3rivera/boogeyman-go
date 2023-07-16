package services

import (
	"github.com/r3rivera/boogeyman/dba"
	"github.com/r3rivera/boogeyman/services/hasher"
)

type Credential struct {
	Email    string
	Password string
}

func CreateUserCredential(email, password string) error {
	i := hasher.Sha256HashItem(password)
	hashItem, _ := i.HashItem()

	creds := dba.NewDDBUserCredential(email, hashItem)
	err := creds.WriteDB()
	return err
}
