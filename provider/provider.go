package provider

import (
	"bytes"
	"github.com/rfsbraz/faker/data"
	"text/template"
	"strings"
	"math/rand"
	"math"
)

type Provider struct {
	Locale string
	Path string
}

func (provider *Provider) Execute(category string) string {
	//We create a template from the required category
	tmpl, err := template.New("t").Parse(data.Load(category, provider.Locale, provider.Path))

	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	
	// Execute the first template
	err = tmpl.Execute(&buffer, provider)

	if err != nil {
		panic(err)
	}

	// We iterate trough the string, and process all remaining templates
	//todo: better template detection
	for ; strings.Contains(buffer.String(), "{{"); tmpl, err = template.New("t").Parse(buffer.String()){
		buffer.Reset()

		// Execute the new template
		err = tmpl.Execute(&buffer, provider)

		if err != nil {
			panic(err)
		}
	}

	return buffer.String()
}

func (provider *Provider) Load(category string) string {
	return provider.LoadFrom(provider.Path, category)
}

func (provider *Provider) LoadFrom(path, category string) string {
	return data.Load(category, provider.Locale, path)
}

func (provider *Provider) Number(length int) int {
    return rand.Intn(int(math.Pow(10, float64(length))))
}

func (provider *Provider) Digit() int {
    return rand.Intn(10)
}

type Person struct {
	Provider
}

type Address struct {
	Provider
}

type Barcode struct {
	Provider
}

type Color struct {
	Provider
	allColors map[string]string
	safeColors []string
}