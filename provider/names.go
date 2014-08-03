package provider

import (
	"bytes"
	"github.com/rfsbraz/faker/data"
	"text/template"
)

type Names struct {
	Locale string
}

func NewNames(locale string) *Names {
	return &Names{Locale: locale}
}

func (names *Names) FirstName() string {
	return data.Load("first_names", names.Locale)
}

func (names *Names) LastName() string {
	return data.Load("last_names", names.Locale)
}

func (names *Names) Title() string {
	return data.Load("titles", names.Locale)
}

func (names *Names) Name() string {
	tmpl, err := template.New("Name").Parse(data.Load("names", names.Locale))
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, names)
	if err != nil {
		panic(err)
	}

	return buffer.String()
}
