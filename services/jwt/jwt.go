package jwt

import "github.com/golang-jwt/jwt"

type Signer interface {
	GenerateJWS() (string, error)
}

type Verifier interface {
	VerifyJWS() error
}

type ClaimExtractor interface {
	ExtractClaims() (jwt.MapClaims, error)
}
