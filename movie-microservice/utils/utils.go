/*
 * @File: utils.utils.go
 * @Description: Reusable stuffs for services
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package utils

import (
	"time"

	"../common"
	jwt_lib "github.com/dgrijalva/jwt-go"
)

// SdtClaims defines the custom claims
type SdtClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt_lib.StandardClaims
}

// GenerateJWT generates token from the given information
func GenerateJWT(name string, role string) (string, error) {
	claims := SdtClaims{
		name,
		role,
		jwt_lib.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    common.Config.Issuer,
		},
	}

	token := jwt_lib.NewWithClaims(jwt_lib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(common.Config.JwtSecretPassword))

	return tokenString, err
}
