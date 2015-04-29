package provider

func NewAddress(locale string) *Address {
	return &Address{Provider{locale, "address"}}
}

func (address *Address) StreetName() string {
	return address.Execute("street_name")
}

func (address *Address) StreetAddress() string {
	return address.Execute("street_address")
}

func (address *Address) Address() string {
	return address.Execute("address")
}

func (address *Address) Postcode() string {
	return address.Execute("postcode")
}

func (address *Address) City() string {
	return address.Execute("city")
}

func (address *Address) Country() string {
	return address.Execute("country")
}