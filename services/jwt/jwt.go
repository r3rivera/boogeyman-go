package jwt

type Signer interface {
	GenerateJWS() (string, error)
}

type Verifier interface {
	VerifyJWS() error
}
