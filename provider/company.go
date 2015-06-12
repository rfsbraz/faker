package provider

func NewCompany(locale, fallback_locale string) *Company {
	return &Company{Provider{locale, fallback_locale, "company"}}
}

func (company *Company) Name() string {
	return company.Execute("company")
}

func (company *Company) CatchPhrase() string {
	return company.Execute("catch_phrase")
}
