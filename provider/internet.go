package provider

func NewInternet(locale string) *Internet {
	return &Internet{Provider{locale, "internet"}}
}
