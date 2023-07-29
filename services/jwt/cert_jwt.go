package jwt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

type CertFile struct {
	sub      string
	fileName string
	claims   map[string]string
}

type JWSToken struct {
	token    string
	fileName string
}

type Claim struct {
	jwt.MapClaims
}

func NewCertFile(sub, file string, claims map[string]string) *CertFile {
	return &CertFile{
		sub:      sub,
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
	cBytes, err := ioutil.ReadFile(f.fileName)
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
		"iss": "r2-rivera.com",
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

	pBytes, err := ioutil.ReadFile(t.fileName)
	if err != nil {
		return errors.New("Unable to read file!")
	}

	pub, err := jwt.ParseRSAPublicKeyFromPEM(pBytes)
	if err != nil {
		return err
	}

	claim := &Claim{}
	token, err := jwt.ParseWithClaims(t.token, claim, func(token *jwt.Token) (interface{}, error) {
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
