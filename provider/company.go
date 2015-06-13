package provider

func NewCompany(faker *Faker) *Company {
	return &Company{faker, "company"}
}

func (company *Company) Name() string {
	return company.Execute("company", company.FolderName)
}

func (company *Company) CatchPhrase() string {
	return company.Execute("catch_phrase", company.FolderName)
}
