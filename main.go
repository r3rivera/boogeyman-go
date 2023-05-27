package main

import (
	"fmt"

	"github.com/r3rivera/boogeyman/b_qrcode"
)

func main() {

	p := b_qrcode.Person{
		FirstName: "Rommel Ryan",
		LastName:  "Rivera",
		Street:    "1420 W MCDERMOTT DR., #417",
		City:      "Allen",
		Zip:       "75013",
		State:     "TEXAS",
		Country:   "USA",
		Phone:     "+1 (650) 773-1121",
	}
	fmt.Println(p.Stringify())

	c := b_qrcode.ContactInfo{
		Email: "rr.rivs@gmail.com",
	}

	fmt.Println(c.Stringify())
	b_qrcode.GenerateQRCode(&c)

}
