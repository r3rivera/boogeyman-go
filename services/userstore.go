package services

type UserRegistrar interface {
	Register() error
}

func RegisterUser(user UserRegistrar) error {
	err := user.Register()
	if err != nil {
		return err
	}
	return nil
}
