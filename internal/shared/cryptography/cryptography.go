package cryptography

import "golang.org/x/crypto/bcrypt"

type Hash struct{}

func NewCryptographyService() *Hash {
	return &Hash{}
}

func (h *Hash) GenerateHash(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), 6)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *Hash) VerifyHash(hash, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))

	return err == nil
}
