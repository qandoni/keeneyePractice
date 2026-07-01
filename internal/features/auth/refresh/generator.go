package auth_refresh

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) Generate() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (g *Generator) Hash(
	token string,
) string {
	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}
