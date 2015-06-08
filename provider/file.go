package provider

func NewFile(locale string) *File {
	return &File{Provider{locale, "file"}}
}
