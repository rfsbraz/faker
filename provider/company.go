package provider

func NewCompany(locale string) *Company {
	return &Company{Provider{locale, "company"}}
}

func (company *Company) Name() string {
	return company.Execute("company")
}

func (company *Company) CatchPhrase() string {
	return company.Execute("catch_phrase")
}
