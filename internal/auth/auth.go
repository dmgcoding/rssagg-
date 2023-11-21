package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization info not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed authorzation info!")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed authorzation first part")
	}

	return vals[1], nil
}
