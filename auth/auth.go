package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("rahasia")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     uint   `json:"role"`
	jwt.StandardClaims
}

func GenerateJTW(email, username string, role uint) (tokenString string, err error) {
	expTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(JwtKey)

	return
}

func ValidateToken(signedToken string) (role uint, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(JwtKey), nil
		},
	)

	if err != nil {
		return
	}
	//jika claim gagal

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("couldnot parse claim with token")
		return
	}

	//jika claim expired
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	// = claims.Email
	role = claims.Role
	return
}
