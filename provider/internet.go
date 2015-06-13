package provider

import (
	"strings"
)

func NewInternet(faker *Faker) *Internet {
	return &Internet{faker, "internet"}
}

func (internet *Internet) Username() string {
	return strings.ToLower(internet.Execute("username", internet.FolderName))
}

func (internet *Internet) Email() string {
	return strings.ToLower(internet.Execute("email", internet.FolderName))
}

func (internet *Internet) URL() string {
	return strings.ToLower(internet.Execute("url", internet.FolderName))
}

func (internet *Internet) URI() string {
	return strings.ToLower(internet.Execute("uri", internet.FolderName))
}

func (internet *Internet) Image(width, height string) string {
	image := internet.Execute("image_placeholder_service", internet.FolderName)
	image = strings.Replace(image, "{Width}", width, -1)
	image = strings.Replace(image, "{Height}", height, -1)
	return image
}
