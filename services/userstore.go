package services

type UserRegistrar interface {
	Register() (bool, error)
}
