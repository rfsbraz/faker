package main

import (
	"github.com/rfsbraz/faker/provider"
	"log"
)

const DEFAULT_LOCALE = "en_US"

type Faker struct {
	Person     *provider.Person
	Address    *provider.Address
	Barcode    *provider.Barcode
	Color      *provider.Color
	Company    *provider.Company
	CreditCard *provider.CreditCard
	Currency   *provider.Currency
	File       *provider.File
}

func NewFaker(locale string) *Faker {
	faker := &Faker{}

	faker.Person = provider.NewPerson(locale)
	faker.Address = provider.NewAddress(locale)
	faker.Barcode = provider.NewBarcode(locale)
	faker.Color = provider.NewColor(locale)
	faker.Company = provider.NewCompany(locale)
	faker.CreditCard = provider.NewCreditCard(locale)
	faker.Currency = provider.NewCurrency(locale)
	faker.File = provider.NewFile(locale)

	return faker
}

func main() {
	faker := NewFaker("pt_PT")
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
}
