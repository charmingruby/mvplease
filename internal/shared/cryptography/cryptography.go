package cryptography

type Hash struct{}

func NewCryptographyService() *Hash {
	return &Hash{}
}

func (h *Hash) GenerateHash(value string) (string, error) {
	return "", nil
}

func (h *Hash) VerifyHash(hash, value string) bool {
	return false
}
