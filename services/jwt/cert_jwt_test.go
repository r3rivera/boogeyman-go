package jwt

import "testing"

func Test_SignJwt(t *testing.T) {

	cert := NewCertFile(
		"test@test.com",
		"/Users/r2devops/Devops/projects/golang/boogeyman-go/private_key.pem")
	jws, err := cert.GenerateJWT()
	if err != nil {
		t.Logf("Failed because :: %s", err.Error())
		t.Fail()
	} else {
		t.Logf("\n\nJWS is %s \n\n", jws)
	}
}
