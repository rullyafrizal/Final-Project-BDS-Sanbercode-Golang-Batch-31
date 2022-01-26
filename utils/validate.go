package utils

import "net/url"

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