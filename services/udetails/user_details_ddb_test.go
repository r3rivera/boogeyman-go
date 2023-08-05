package udetails

import (
	"testing"
	"time"
)

func Test_GetCredenti(t *testing.T) {

	dob, err := time.Parse("2006-01-02", "2013-10-18")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	info := UserDetail{
		firstname: "Junit",
		lastname:  "Tester",
		dob:       dob,
	}

	address := UserAddress{
		street1: "123 Main Street",
		street2: "Suite 30",
		city:    "San Antonio",
		state:   "TX",
		zip:     "78238",
	}

	creds := NewUserDetail("junit@test.com", info, address)
	err = creds.WriteDB()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

}
