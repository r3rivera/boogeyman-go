package dba

import "testing"

func Test_GetCredenti(t *testing.T) {

	creds := NewDDBUserCredential("reizen@test.com", "test123")
	_, err := creds.ReadDB()
	if err != nil {
		t.Fail()
	}

}
