package libraries

import (
	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var mySigningKey = []byte("secret-rahasia")

func ValidateToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*MyCustomClaims)
	return token.Valid, claims.Username
}

func ClaimToken(email, username, role string) (string, error) {
	claims := MyCustomClaims{
		email,
		username,
		role,
		jwt.StandardClaims{
			Subject: "encode-token",
			Id:      username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(mySigningKey)
}
