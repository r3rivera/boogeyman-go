package main

import (
	"fmt"

	bhash "github.com/r3rivera/boogeyman/b_hash"
	bqrcode "github.com/r3rivera/boogeyman/b_qrcode"
)

type Person struct {
	FirstName string
	LastName  string
	Street    string
	City      string
	State     string
	Zip       string
	Country   string
	Phone     string
}

type ContactInfo struct {
	Email string
}

func (c *ContactInfo) Stringify() string {
	return fmt.Sprintf("mailto:%s", c.Email)
}

func (c *ContactInfo) DataToHash() string {
	return fmt.Sprintf("|%s|", c.Email)
}

func (p *Person) Stringify() string {
	return fmt.Sprintf("\n\n%s %s\n\n%s\n%s %s, %s %s\n\n %s", p.FirstName, p.LastName, p.Street,
		p.City, p.Zip, p.State, p.Country, p.Phone)
}

func (p *Person) DataToHash() string {
	return fmt.Sprintf("|%s %s|%s %s %s %s %s|%s", p.FirstName, p.LastName, p.Street,
		p.City, p.Zip, p.State, p.Country, p.Phone)
}

func main() {

	p := Person{
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

	c := ContactInfo{
		Email: "",
	}

	fmt.Println(c.Stringify())
	bqrcode.GenerateQRCode(&c)
	fmt.Println(bhash.GenerateHash(&c))

}
