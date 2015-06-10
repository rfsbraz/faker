package provider

import (
	"strings"
)

func NewInternet(locale string) *Internet {
	return &Internet{Provider{locale, "internet"}}
}

func (internet *Internet) Username() string {
	return internet.Execute("username")
}

func (internet *Internet) Email() string {
	return internet.Execute("email")
}

func (internet *Internet) URL() string {
	return internet.Execute("url")
}

func (internet *Internet) URI() string {
	return internet.Execute("uri")
}

func (internet *Internet) Image(width, height string) string {
	image := internet.Execute("image_placeholder_service")
	image = strings.Replace(image, "{Width}", width, -1)
	image = strings.Replace(image, "{Height}", height, -1)
	return image
}
