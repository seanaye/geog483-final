package jwt

import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/seanaye/geog483-final/server/pkg/random"
)


func init() {
	secret := random.RandString(24)
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
