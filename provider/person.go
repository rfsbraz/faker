package provider

func NewPerson(locale string) *Person {
	return &Person{Provider{locale}}
}

func (person *Person) FirstName() string {
	return person.Execute("first_names")
}

func (person *Person) LastName() string {
	return person.Execute("last_names")
}

func (person *Person) Title() string {
	return person.Execute("titles")
}

func (person *Person) Name() string {
	return person.Execute("names")
}