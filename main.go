package main

import (
	"github.com/rfsbraz/faker/provider"
	"log"
)

func main() {
	faker := provider.NewFaker("pt_PT", "en_US")
	log.Println(faker.Person.Title())
	log.Println(faker.Person.Name())
	log.Println(faker.Address.Address())
	log.Println(faker.Barcode.EAN8())
	log.Println(faker.Barcode.EAN13())
	log.Println(faker.Color.ColorName())
	log.Println(faker.Color.HexForColorName(faker.Color.ColorName()))
	log.Println(faker.Color.SafeColorName())
	log.Println(faker.Color.HexColor())
	log.Println(faker.Color.RGBColorList())
	log.Println(faker.Color.RGBColor())
	log.Println(faker.Color.RGBCSSColor())
	log.Println(faker.Company.Name())
	log.Println(faker.Company.CatchPhrase())
	log.Println(faker.CreditCard.Full("amex"))
	log.Println(faker.Currency.Code())
	log.Println(faker.File.MimeType(""))
	log.Println(faker.File.Name(""))
	log.Println(faker.File.Extension(""))
	log.Println(faker.Internet.Username())
	log.Println(faker.Internet.Email())
	log.Println(faker.Internet.FreeEmail())
	log.Println(faker.Internet.URL())
	log.Println(faker.Internet.URI())
	log.Println(faker.Internet.Image("500", "500"))
}
