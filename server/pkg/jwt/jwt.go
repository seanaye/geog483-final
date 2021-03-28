package jwt

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/seanaye/geog483-final/server/pkg/random"
)


func init() {
	secret := random.RandString(24)
	fmt.Print("Starting server with secret ", secret, "\n")
	os.Setenv("JWT_ACCESS", secret)
}

func CreateToken(id string, name string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["name"] = name
	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_ACCESS")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateToken(str string) (*jwt.Token error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(str, claims, func(jwt *jwt.Token) (interface{} error) {
		return jwt, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.valid {
		return nil, errors.New("Token is invalid")
	}

	return token, nil
}
