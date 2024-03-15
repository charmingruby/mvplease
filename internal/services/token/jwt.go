package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTService struct {
	issuer    string
	secretKey string
}

func NewJWTService(secretKey, issuer string) *JWTService {
	return &JWTService{
		issuer:    issuer,
		secretKey: secretKey,
	}
}

type JWTClaim struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func (j *JWTService) GenerateToken(accountID uuid.UUID, role string) (string, error) {
	tokenDuration := time.Duration(time.Minute * 60 * 24 * 7) //7 days

	claims := &JWTClaim{
		role,
		jwt.StandardClaims{
			Subject:   accountID.String(),
			Issuer:    j.issuer,
			ExpiresAt: time.Now().Local().Add(tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", nil
	}

	return tokenStr, nil
}
