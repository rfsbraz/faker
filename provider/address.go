package provider

func NewAddress(faker *Faker) *Address {
	return &Address{faker, "address"}
}

func (address *Address) StreetName() string {
	return address.Execute("street_name", address.FolderName)
}

func (address *Address) StreetAddress() string {
	return address.Execute("street_address", address.FolderName)
}

func (address *Address) Address() string {
	return address.Execute("address", address.FolderName)
}

func (address *Address) Postcode() string {
	return address.Execute("postcode", address.FolderName)
}

func (address *Address) City() string {
	return address.Execute("city", address.FolderName)
}

func (address *Address) Country() string {
	return address.Execute("country", address.FolderName)
}
