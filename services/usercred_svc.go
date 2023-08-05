package services

import (
	"log"

	"github.com/r3rivera/boogeyman/services/hasher"
	"github.com/r3rivera/boogeyman/services/ucreds"
)

type Credential struct {
	Email    string
	Password string
}

func CreateUserCredential(email, password string) error {
	i := hasher.NewBCrypt(password)
	hashItem, _ := i.HashItem()

	creds := ucreds.NewDDBUserCredential(email, hashItem)
	err := creds.WriteDB()
	return err
}

func VerifyUserCredential(email, password string) (bool, error) {
	output := ucreds.NewDDBUserCredential(email, password)

	hashOut, err := output.ReadDB()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return (hasher.BCryptVerifyItem(hashOut, password)), nil

}
