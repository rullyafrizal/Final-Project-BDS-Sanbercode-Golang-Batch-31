package utils

import (
	"net/mail"
	"net/url"
)

func IsValidUrl(v string) bool {
	_, err := url.ParseRequestURI(v)
	if err != nil {
		return false
	}

	u, err := url.Parse(v)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func IsValidEmail(v string) bool {
	_, err := mail.ParseAddress(v)
	
    return err == nil
}