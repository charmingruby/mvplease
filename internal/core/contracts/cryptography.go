package contracts

type CryptographyContract interface {
	GenerateHash(value string) (string, error)
	VerifyHash(hash, value string) bool
}
