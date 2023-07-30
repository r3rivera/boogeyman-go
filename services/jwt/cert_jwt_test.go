package jwt

import (
	"testing"
	"time"
)

func Test_SignJwt(t *testing.T) {

	custom := map[string]string{
		"role":  "tester",
		"scope": "testing",
	}

	cert := NewCertFile(
		"test@test.com",
		"localhost.com",
		"/var/local/private1_key.pem",
		custom)
	jws, err := cert.GenerateJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	} else {
		t.Logf("\n\nJWS is %s \n\n", jws)
	}
}

func Test_VerifyJws(t *testing.T) {

	cert := NewCertFile(
		"test@test.com",
		"localhost.com",
		"/var/local/private1_key.pem",
		nil)
	jws, err := cert.GenerateJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

	pubKey := "/var/local/public1_key.pem"
	verifier := NewTokenVerifier(jws, pubKey)

	err = verifier.VerifyJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

}

func Test_ExtractClaim(t *testing.T) {

	customClaims := map[string]string{
		"role": "user",
		"app":  "app1",
	}

	cert := NewCertFile(
		"test@test.com",
		"localhost.com",
		"/var/local/private1_key.pem",
		customClaims)
	jws, err := cert.GenerateJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

	pubKey := "/var/local/public1_key.pem"
	verifier := NewTokenVerifier(jws, pubKey)

	claim, err := verifier.ExtractClaims()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}
	if !claim.VerifyIssuer("localhost.com", true) {
		t.Log("Failed because :: Invalid Issuer")
		t.Fail()
	}
	t.Log(claim)
}

func Test_ValidateClaim(t *testing.T) {

	customClaims := map[string]string{
		"role": "user",
		"app":  "app1",
	}

	const ISS = "localhost.com"
	cert := NewCertFile(
		"test@test.com",
		ISS,
		"/var/local/private1_key.pem",
		customClaims)
	jws, err := cert.GenerateJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

	pubKey := "/var/local/public1_key.pem"
	verifier := NewTokenVerifier(jws, pubKey)

	time.Sleep(time.Duration(5) * time.Second)
	isValid := ValidateClaims(ISS, verifier)
	if !isValid {
		t.Fail()
	}

}
