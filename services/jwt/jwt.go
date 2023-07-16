package jwt

type Signer interface {
	GenerateJWT() (string, error)
}
