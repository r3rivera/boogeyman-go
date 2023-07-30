package jwt

import "testing"

func Test_SignJwt(t *testing.T) {

	custom := map[string]string{
		"role":  "tester",
		"scope": "testing",
	}

	cert := NewCertFile(
		"test@test.com",
		"localhost.com",
		"/Users/r2devops/Devops/projects/golang/boogeyman-go/private1_key.pem",
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
		"/Users/r2devops/Devops/projects/golang/boogeyman-go/private1_key.pem",
		nil)
	jws, err := cert.GenerateJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

	pubKey := "/Users/r2devops/Devops/projects/golang/boogeyman-go/public1_key.pem"
	verifier := NewTokenVerifier(jws, pubKey)

	err = verifier.VerifyJWS()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	}

}
