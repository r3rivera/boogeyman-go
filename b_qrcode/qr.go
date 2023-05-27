package bqrcode

import (
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
