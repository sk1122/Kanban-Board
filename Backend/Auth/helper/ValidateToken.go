package helper

import (
	"strings"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	tokenArr := strings.Split(bearToken, " ")

	if len(tokenArr) == 2 {
		return tokenArr[1]
	}

	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	ACCESS_SECRET := "jdnfksdmfksd"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     	}
     	return []byte(ACCESS_SECRET), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) (interface{}, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	return username, nil
}