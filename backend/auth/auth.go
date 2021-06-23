package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var (
	err = godotenv.Load()
	key = os.Getenv("SECRET_KEY")
)

type Service interface {
	GenerateToken(PetaniID int) (string, error)
	GenerateTokenInvestor(InvestorID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
	GenerateTokenKpetani(KPetaniID int) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(PetaniID int) (string, error) {
	claim := jwt.MapClaims{
		"petani_id": PetaniID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}

func (s *jwtService) GenerateTokenInvestor(InvestorID int) (string, error) {
	claim := jwt.MapClaims{
		"investor_id": InvestorID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}

func (s *jwtService) GenerateTokenKpetani(KPetaniID int) (string, error) {
	claim := jwt.MapClaims{
		"kpetani_id": KPetaniID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(key))

	if err != nil {
		return signedToken, nil
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(encodedToken *jwt.Token) (interface{}, error) {
		_, ok := encodedToken.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
