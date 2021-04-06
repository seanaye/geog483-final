package jwt

import (
	"fmt"
	"os"
	"time"
	"log"

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

func ValidateClaims(tokenStr string) (jwt.MapClaims, bool) {
		hmacSecretString := os.Getenv("JWT_ACCESS")
		hmacSecret := []byte(hmacSecretString)
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				 // check token signing method etc
				 return hmacSecret, nil
		})

		if err != nil {
				return nil, false
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				return claims, true
		} else {
				log.Printf("Invalid JWT Token")
				return nil, false
		}
}

