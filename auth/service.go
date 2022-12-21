package auth

import "github.com/golang-jwt/jwt/v4"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type JWTService struct {
}

func NewService() *JWTService {
	return &JWTService{}
}

var SECRETKEY = []byte("Budianto_Secret_Key")

func (s *JWTService) GenerateToken(userID int) (string, error) {
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(SECRETKEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
