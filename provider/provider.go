package provider

import (
	"bytes"
	"github.com/rfsbraz/faker/data"
	"text/template"
	"strings"
)

type Provider struct {
	Locale string
}

func (provider *Provider) Execute(category string) string {

	//We create a template from the required category
	tmpl, err := template.New("t").Parse(data.Load(category, provider.Locale))

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
	return data.Load(category, provider.Locale)
}

type Person struct {
	Provider
}