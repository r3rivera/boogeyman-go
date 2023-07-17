package jwt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

type CertFile struct {
	sub      string
	fileName string
	claims   map[string]string
}

func NewCertFile(sub, file string, claims map[string]string) *CertFile {
	return &CertFile{
		sub:      sub,
		fileName: file,
		claims:   claims,
	}
}

func (f *CertFile) GenerateJWT() (string, error) {
	log.Printf("PATH is %s", f.fileName)
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