package ucreds

import "testing"

func Test_GetCredenti(t *testing.T) {

	creds := NewDDBUserCredential("junit@test.com", "test123")
	err := creds.WriteDB()
	if err != nil {
		t.Fail()
	}

	hash, err := creds.ReadDB()
	if err != nil {
		t.Fail()
	}
	if hash == "" {
		t.Fail()
	}

	err1 := creds.DeleteDB()
	if err1 != nil {
		t.Fail()
	}

}
