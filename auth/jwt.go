package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSecretKey = []byte("rahul123")
var refreshSecretKey = []byte("rahul123")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {

	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "dasboardapi",
			Subject:   username,
			ID:        "1",
			Audience:  []string{"Not_applicable"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(tokenSecretKey)
	if err != nil {
		return "", errors.New("token creation failed")
	}
	fmt.Printf("%v %v", ss, err)
	return ss, nil
}

func GenerateRefreshToken(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "dasboardapi",
			Subject:   username,
			ID:        "1",
			Audience:  []string{"Not_applicable"},
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedRefreshToken, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return "", errors.New("refresh token creation failed")
	}
	fmt.Printf("%v %v", signedRefreshToken, err)
	return signedRefreshToken, nil

}

// func TokenMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("heyyyyyyy token is valid")
// 		params := r.Header
// 		fmt.Println("params issss", params)
// 		next.ServeHTTP(w, r)
// 		fmt.Println("donennnn")

// 	})

// }
