package helper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string) (string, string, error) {
	var aterr error
	var rterr error

	ACCESS_SECRET := "jdnfksdmfksd"

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	access_token, aterr := at.SignedString([]byte(ACCESS_SECRET))

	if aterr != nil {
		return "", "", aterr
	}

	REFRESH_SECRET := "mcmvmkmsdnfsdmfdsjf"

	rtClaims := jwt.MapClaims{}
	rtClaims["authorized"] = true
	rtClaims["username"] = username
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	refresh_token, rterr := rt.SignedString([]byte(REFRESH_SECRET))

	if rterr != nil {
		return "", "", rterr
	}

	return access_token, refresh_token, aterr
}