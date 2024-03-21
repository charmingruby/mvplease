package cryptography

import "golang.org/x/crypto/bcrypt"

func NewCryptographyService() *CryptographyService {
	return &CryptographyService{}
}

type CryptographyService struct{}

func (h *CryptographyService) GenerateHash(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 6)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *CryptographyService) VerifyHash(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
