package gogopro

import (
	"strings"
)

func cleanURL(baseURL string) (string, error) {
	if strings.HasSuffix(baseURL, "/") {
		return baseURL, nil
	}
	return baseURL + "/", nil
}
