package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

type CertFile struct {
	sub      string
	iss      string
	fileName string
	claims   map[string]string
}

type JWSToken struct {
	token    string
	fileName string
}

func NewCertFile(sub, issuer, file string, claims map[string]string) *CertFile {
	return &CertFile{
		sub:      sub,
		iss:      issuer,
		fileName: file,
		claims:   claims,
	}
}

func NewTokenVerifier(token, pathName string) *JWSToken {
	return &JWSToken{
		token:    token,
		fileName: pathName,
	}
}

func (f *CertFile) GenerateJWS() (string, error) {
	cBytes, err := os.ReadFile(f.fileName)
	if err != nil {
		return "", errors.New("Unable to read file!")
	}

	cert, err := ssh.ParseRawPrivateKey(cBytes)
	if err != nil {
		msg := fmt.Sprintf("Error::Parse Found :: %v", err.Error())
		return "", errors.New(msg)
	}

	claims := jwt.MapClaims{
		"sub": f.sub,
		"iss": f.iss,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	if f.claims != nil {
		for k, v := range f.claims {
			claims[k] = v
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jws, err := token.SignedString(cert)
	if err != nil {
		msg := fmt.Sprintf("Error::Sign Found :: %v", err.Error())
		return "", errors.New(msg)
	}
	return jws, nil
}

func (t *JWSToken) VerifyJWS() error {

	pBytes, err := os.ReadFile(t.fileName)
	if err != nil {
		return errors.New("Unable to read file!")
	}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pBytes)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(t.token, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New("Invalid Token Signature")
		}
		return err
	}
	if !token.Valid {
		return errors.New("Token Not Valid")
	}
	return nil
}

func (t *JWSToken) ExtractClaims() (jwt.MapClaims, error) {

	pBytes, err := os.ReadFile(t.fileName)
	if err != nil {
		return nil, errors.New("Unable to read file!")
	}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pBytes)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(t.token, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid Token Signature")
		}
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token Not Valid")
	}

}
