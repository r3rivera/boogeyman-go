package services

import (
	"testing"
)

func Test_hashing(t *testing.T) {

	hash, err := HashPassword("MyTestPasswOrd$")
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)

	notSame := VerifyPassword("MyTestPasswOrd$", hash)
	if notSame {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
}
