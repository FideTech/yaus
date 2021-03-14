package utils

import (
	"net/url"
)

//IsValidUrl checks if the provided value is a valid url with a valid scheme and host.
//This is only a bare basic check to catch basic errors and not complex ones
func IsValidUrl(val string) bool {
	u, err := url.Parse(val)
	return err == nil && u.Scheme != "" && u.Host != ""
}
