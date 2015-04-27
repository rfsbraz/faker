package main

import (
	"github.com/rfsbraz/faker/provider"
	"log"
)

const DEFAULT_LOCALE = "en_US"

type Faker struct {
	Person *provider.Person
}

func NewFaker(locale string) *Faker {
	faker := &Faker{}

	faker.Person = provider.NewPerson(locale)

	return faker
}

func main() {
	faker := NewFaker("pt_PT")
	log.Println(faker.Person.Name())
}
