package helper

import "net/url"

// Making sure passed string follows the correct URL scheme
func IsValidUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}
