package main

import (
	"github.com/rfsbraz/faker/provider"
	"log"
)

const DEFAULT_LOCALE = "en_US"

type Faker struct {
	Names *provider.Names
}

func NewFaker(locale string) *Faker {
	faker := &Faker{}

	faker.Names = provider.NewNames(locale)

	return faker
}

func main() {
	faker := NewFaker("pt_PT")
	log.Println(faker.Names.Name())
}
