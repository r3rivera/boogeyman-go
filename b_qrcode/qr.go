package b_qrcode

import (
	"fmt"

	"github.com/skip2/go-qrcode"
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

type Stringified interface {
	Stringify() string
}

func (c *ContactInfo) Stringify() string {
	return fmt.Sprintf("mailto:%s", c.Email)
}

func (p *Person) Stringify() string {
	return fmt.Sprintf("\n\n%s %s\n\n%s\n%s %s, %s %s\n\n %s", p.FirstName, p.LastName, p.Street,
		p.City, p.Zip, p.State, p.Country, p.Phone)
}

func GenerateQRCode(item Stringified) error {

	data := item.Stringify()
	qrCode, err := qrcode.New(data, qrcode.High)
	if err != nil {
		return err
	}

	err = qrCode.WriteFile(200, "r2.png")
	if err != nil {
		return err
	}

	return nil
}
