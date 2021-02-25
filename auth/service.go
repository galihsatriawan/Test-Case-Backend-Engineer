package auth

import (
	"errors"
	"fmt"
	"os"
	"test-case-backend/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}
func getSecretKey() []byte {
	err := godotenv.Load("backend.env")
	helper.FailOnError(err, "Error when trying to env")
	return []byte(os.Getenv("MY_SECRET_KEY"))
}
func (s *jwtService) GenerateToken(userId int) (string, error) {

	var my_secret_key = getSecretKey()
	// Insert data payload
	payload := jwt.MapClaims{}
	payload["user_id"] = userId

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Signature
	fmt.Println(my_secret_key)
	signedToken, err := token.SignedString(my_secret_key)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	var my_secret_key = getSecretKey()

	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if comma, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println(comma)
			return encodedToken, errors.New("Invalid Token")
		}
		return []byte(my_secret_key), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
