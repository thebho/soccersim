package util

import (
	"net/url"
)

// ParseURL reads query params from a URL
func ParseURL(queryString, urlArg string) (string, error) {
	u, err := url.Parse(urlArg)
	if err != nil {
		return "", err
	}
	q := u.Query()
	return q.Get(queryString), nil
}
