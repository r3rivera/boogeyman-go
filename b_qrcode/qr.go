package b_qrcode

import (
	"fmt"
	"log"

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

func GenerateQRCode(item Person) error {

	data := fmt.Sprintf("\n\n%s %s\n\n%s\n%s %s, %s %s\n\n %s", item.FirstName, item.LastName, item.Street,
		item.City, item.Zip, item.State, item.Country, item.Phone)
	log.Println(data)

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
