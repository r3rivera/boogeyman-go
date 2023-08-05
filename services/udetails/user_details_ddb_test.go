package udetails

import (
	"fmt"
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
		Firstname: "Junit",
		Lastname:  "Tester",
		Dob:       dob,
	}

	address := UserAddress{
		Street1: "123 Main Street",
		Street2: "Suite 30",
		City:    "San Antonio",
		State:   "TX",
		Zip:     "78238",
	}

	creds := NewUserDetail("junit@test.com", info, address)
	err = creds.WriteDB()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	user, err := creds.ReadDB()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	fmt.Printf("Response is %v", user)

	err = creds.DeleteDB()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

}
