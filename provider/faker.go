package provider

import (
	"bytes"
	"fmt"
	"github.com/rfsbraz/faker/data"
	"math/rand"
	"strconv"
	"strings"
	"text/template"
)

type Faker struct {
	Locale         string
	FallbackLocale string

	Person     *Person
	Address    *Address
	Barcode    *Barcode
	Color      *Color
	Company    *Company
	CreditCard *CreditCard
	Currency   *Currency
	File       *File
	Internet   *Internet
}

func NewFaker(locale, fallback_locale string) *Faker {
	faker := &Faker{Locale: locale, FallbackLocale: fallback_locale}

	faker.Person = NewPerson(faker)
	faker.Address = NewAddress(faker)
	faker.Barcode = NewBarcode(faker)
	faker.Color = NewColor(faker)
	faker.Company = NewCompany(faker)
	faker.CreditCard = NewCreditCard(faker)
	faker.Currency = NewCurrency(faker)
	faker.File = NewFile(faker)
	faker.Internet = NewInternet(faker)

	return faker
}

func (faker *Faker) Execute(category, path string) string {
	//We create a template from the required category
	tmpl, err := template.New("t").Parse(data.Load(category, faker.Locale, faker.FallbackLocale, path))

	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer

	// Execute the first template
	err = tmpl.Execute(&buffer, faker)

	if err != nil {
		panic(err)
	}

	// We iterate trough the string, and process all remaining templates
	//todo: better template detection
	for ; strings.Contains(buffer.String(), "{{"); tmpl, err = template.New("t").Parse(buffer.String()) {
		buffer.Reset()

		// Execute the new template
		err = tmpl.Execute(&buffer, faker)

		if err != nil {
			panic(err)
		}
	}

	return buffer.String()
}

func (faker *Faker) LoadFrom(path, category string) string {
	return data.Load(category, faker.Locale, faker.FallbackLocale, path)
}

func (faker *Faker) RandomNumberOfSize(length int) int {
	number := ""
	for i := 0; i < length; i++ {
		digit := faker.RandomDigit()
		number = fmt.Sprintf("%v%v", number, digit)
		if i == 0 && digit == 0 {
			i = -1
		}
	}
	val, _ := strconv.Atoi(number)
	return val
}

func (faker *Faker) RandomNumber(max int) int {
	return rand.Intn(max)
}

func (faker *Faker) RandomDigit() int {
	return rand.Intn(10)
}

type Person struct {
	*Faker
	FolderName string
}

type Address struct {
	*Faker
	FolderName string
}

type Barcode struct {
	*Faker
	FolderName string
}

type Color struct {
	*Faker
	FolderName string
}

type Company struct {
	*Faker
	FolderName string
}

type CreditCard struct {
	*Faker
	FolderName string
}

type Currency struct {
	*Faker
	FolderName string
}

type File struct {
	*Faker
	FolderName string
}

type Internet struct {
	*Faker
	FolderName string
}
