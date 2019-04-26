package libraries

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Uid string `json:"uid"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret-shejek")

func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Uid
}

func ClaimToken(uid string) (string, error) {
	claims := MyCustomClaims{
		uid,
		jwt.StandardClaims{
			Subject: "encode-token",
			Id:      uid,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(mySigningKey)
}
