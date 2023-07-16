package hasher

import (
	"fmt"
	"testing"
)

func Test_Hash(t *testing.T) {
	item := Sha256HashItem("SomeTestValue")
	value, err := item.HashItem()
	fmt.Println("SHA256 Value is " + value)
	if err != nil {
		t.Fail()
	}

	item2 := Sha256HashItem("SomeTestValue")
	value2, err := item2.HashItem()
	if err != nil {
		t.Fail()
	}
	if value != value2 {
		t.Log("Invalid Value")
		t.Fail()
	}

}
