package token

import (
	"fmt"
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

func (j *JWTService) isTokenValid(t *jwt.Token) (interface{}, error) {
	if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, fmt.Errorf("invalid token %v", t)
	}

	return []byte(j.secretKey), nil
}

func (j *JWTService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, j.isTokenValid)

	return err == nil
}

type payload struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
}

func (j *JWTService) RetriveTokenPayload(token string) (*payload, error) {
	t, err := jwt.Parse(token, j.isTokenValid)
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("unable to parse jwt claims")
	}

	fmt.Printf("%v\n", claims)

	payload := &payload{
		AccountID: claims["sub"].(string),
		Role:      claims["role"].(string),
	}

	return payload, err
}
