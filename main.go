package main

import (
	"fmt"

	"github.com/r3rivera/boogeyman/b_qrcode"
)

func main() {

	p := b_qrcode.Person{
		FirstName: "",
		LastName:  "",
		Street:    "",
		City:      "",
		Zip:       "",
		State:     "",
		Country:   "",
		Phone:     "",
	}
	fmt.Println(p.Stringify())

	c := b_qrcode.ContactInfo{
		Email: "",
	}

	fmt.Println(c.Stringify())
	b_qrcode.GenerateQRCode(&c)

}
