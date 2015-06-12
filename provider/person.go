package provider

func NewPerson(locale, fallback_locale string) *Person {
	return &Person{Provider{locale, fallback_locale, "person"}}
}

func (person *Person) FirstName() string {
	return person.Execute("first_name")
}

func (person *Person) LastName() string {
	return person.Execute("last_name")
}

func (person *Person) Title() string {
	return person.Execute("title")
}

func (person *Person) Name() string {
	return person.Execute("name")
}
