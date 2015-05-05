package main

import (
	"github.com/rfsbraz/faker/provider"
	"log"
)

const DEFAULT_LOCALE = "en_US"

type Faker struct {
	Person *provider.Person
	Address *provider.Address
	Barcode *provider.Barcode
}

func NewFaker(locale string) *Faker {
	faker := &Faker{}

	faker.Person = provider.NewPerson(locale)
	faker.Address = provider.NewAddress(locale)
	faker.Barcode = provider.NewBarcode()

	return faker
}

func main() {
	faker := NewFaker("pt_PT")
	log.Println(faker.Person.Title())
	log.Println(faker.Person.Name())
	log.Println(faker.Address.Address())
	log.Println(faker.Barcode.EAN8())
	log.Println(faker.Barcode.EAN13())
}
