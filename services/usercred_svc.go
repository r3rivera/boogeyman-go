package services

import (
	"log"

	"github.com/r3rivera/boogeyman/dba"
	"github.com/r3rivera/boogeyman/services/hasher"
)

type Credential struct {
	Email    string
	Password string
}

func CreateUserCredential(email, password string) error {
	i := hasher.NewBCrypt(password)
	hashItem, _ := i.HashItem()

	creds := dba.NewDDBUserCredential(email, hashItem)
	err := creds.WriteDB()
	return err
}

func VerifyUserCredential(email, password string) (bool, error) {
	output := dba.NewDDBUserCredential(email, password)

	hashOut, err := output.ReadDB()
	log.Printf("\n\n DB Value is %v \n\n", hashOut)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return (hasher.BCryptVerifyItem(hashOut, password)), nil

}
