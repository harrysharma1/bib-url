package helper

import "net/url"

func IsValidUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
