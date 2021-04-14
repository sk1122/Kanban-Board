package middleware

import (
	"net/http"
	"Backend/Auth/helper"

	"fmt"
)

func GetUser(r *http.Request) (string, error) {
	u, errr := helper.TokenValid(r)

	username := fmt.Sprintf("%v", u)

	if errr == nil {
		return username, nil
	} else {
		return "", errr
	}
}