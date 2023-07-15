package services

import "testing"

func Test_BCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashPassword()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)
	same := pass.VerifyPassword(hash)
	t.Logf("Is Same? %v", same)
	if same {
		t.Log("Expected to not fail")
	} else {
		t.Fail()
	}
}

func Test_NegativeBCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashPassword()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)

	pass = BCryptPassword{"MyTestPasswOrd"}
	same := pass.VerifyPassword(hash)
	t.Logf("Is Same? %v", same)
	if !same {
		t.Log("Expected to fail")
	} else {
		t.Fail()
	}
}

func Test_PositiveBCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashPassword()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)

	pass = BCryptPassword{"MyTestPasswOrd$"}
	same := pass.VerifyPassword(hash)
	t.Logf("Is Same? %v", same)
	if same {
		t.Log("Expected to not fail")
	} else {
		t.Fail()
	}
}
