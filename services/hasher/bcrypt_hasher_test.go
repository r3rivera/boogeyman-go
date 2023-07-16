package hasher

import "testing"

func Test_BCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashItem()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)
	same := pass.VerifyItem(hash)
	t.Logf("Is Same? %v", same)
	if same {
		t.Log("Expected to not fail")
	} else {
		t.Fail()
	}
}

func Test_NegativeBCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashItem()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)

	pass = BCryptPassword{"MyTestPasswOrd"}
	same := pass.VerifyItem(hash)
	t.Logf("Is Same? %v", same)
	if !same {
		t.Log("Expected to fail")
	} else {
		t.Fail()
	}
}

func Test_PositiveBCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash, err := pass.HashItem()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash)

	pass = BCryptPassword{"MyTestPasswOrd$"}
	same := pass.VerifyItem(hash)
	t.Logf("Is Same? %v", same)
	if same {
		t.Log("Expected to not fail")
	} else {
		t.Fail()
	}
}
