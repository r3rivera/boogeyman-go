package jwt

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

type Signer interface {
	GenerateJWS() (string, error)
}

type Verifier interface {
	VerifyJWS() error
}

type privateKeyReader interface {
	GetPrivateKey() (*rsa.PrivateKey, error)
}

type publicKeyReader interface {
	GetPublicKey() (*rsa.PublicKey, error)
}

type PrivatePEMKey string
type PublicPEMKey string

func (k PrivatePEMKey) GetPrivateKey() (*rsa.PrivateKey, error) {
	bytes, err := os.ReadFile(string(k))
	if err != nil {
		return nil, err
	}
	key, err := ssh.ParseRawPrivateKey(bytes)
	if err != nil {
		return nil, err
	}
	return key.(*rsa.PrivateKey), nil
}

func (k PublicPEMKey) GetPublicKey() (*rsa.PublicKey, error) {
	bytes, err := os.ReadFile(string(k))
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
