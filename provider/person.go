package provider

func NewPerson(faker *Faker) *Person {
	return &Person{faker, "person"}
}

func (person *Person) FirstName() string {
	return person.Execute("first_name", person.FolderName)
}

func (person *Person) LastName() string {
	return person.Execute("last_name", person.FolderName)
}

func (person *Person) Title() string {
	return person.Execute("title", person.FolderName)
}

func (person *Person) Name() string {
	return person.Execute("name", person.FolderName)
}
