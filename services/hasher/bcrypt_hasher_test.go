package hasher

import "testing"

func Test_NegativeBCrypt(t *testing.T) {

	pass := BCryptPassword{"MyTestPasswOrd$"}
	hash1, err := pass.HashItem()
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	t.Logf("HashPassword is %s", hash1)
	same := BCryptVerifyItem(hash1, "MyTestPasswOrd!")
	if !same {
		t.Log("Expected to fail")
	} else {
		t.Fail()
	}
}

func Test_PositiveBCrypt(t *testing.T) {

	pass1 := NewBCrypt("MyTestPasswOrd$")
	hash1, err := pass1.HashItem()
	t.Logf("Hash1 Password is %s", hash1)
	if err != nil {
		t.Log("Error found with hashing the password")
		t.Fail()
	}
	same := BCryptVerifyItem(hash1, "MyTestPasswOrd$")
	if same {
		t.Log("Expected to not fail")
	} else {
		t.Fail()
	}
}
