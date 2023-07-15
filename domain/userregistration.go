package domain

import "errors"

type UserRegistration struct {
	FirstName string
	LastName  string
	Email     string
}

func (u *UserRegistration) Register() error {

	return errors.New("Unable to register user :: " + u.FirstName)

}
