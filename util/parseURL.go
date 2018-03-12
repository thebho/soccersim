package util

import (
	"net/url"
)

func ParseURL(queryString, urlArg string) (string, error) {
	u, err := url.Parse(urlArg)
	if err != nil {
		return "", err
	}
	q := u.Query()
	return q.Get(queryString), nil
}
