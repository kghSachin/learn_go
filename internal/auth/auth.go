package auth

import (
	"errors"
	"net/http"
	"strings"
)

//getAPIKey extracts the api key from the headers of an HTTP Request
//lets expect the authorization header to contain the api key
// Authorization : Bearer {api_key here }

func GetAPIKey(headers http.Header) (string, error) {
	print("GetAPIKey")
	print(headers.Get("Authorization"))
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization Info found")
	}
	value := strings.Split(val, " ")
	print(len(value))
	if len(value) != 2 {
		return "", errors.New("invalid authorization Info")
	}
	if value[0] != "Bearer" {
		return "", errors.New("use Bearer as string before sending api key Info")
	}
	return value[1], nil
}
