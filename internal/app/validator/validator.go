package validator

import (
	"regexp"
)

const (
	patternUrl      = `^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`
	patternProtocol = `^https?:\/\/(?:www\.)?.+\/`
)

var (
	regUrl      = regexp.MustCompile(patternUrl)
	regProtocol = regexp.MustCompile(patternProtocol)
)

func ValidateUrl(url string) bool {
	return regUrl.MatchString(url)
}

func TrimProtocol(url string) string {
	return regProtocol.ReplaceAllString(url, "")
}
