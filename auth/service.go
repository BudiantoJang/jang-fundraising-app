package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTUsecase struct {
}

func NewUsecase() *JWTUsecase {
	return &JWTUsecase{}
}

var SECRET_KEY = []byte("Budianto_Secret_Key")

func (s *JWTUsecase) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *JWTUsecase) ValidateToken(token string) (*jwt.Token, error) {
	validatedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return validatedToken, nil
}
