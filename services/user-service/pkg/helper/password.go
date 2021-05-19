package helper

import (
	"bytes"
	"encoding/base64"
	"hash"
	"math/rand"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// Password model
type Password struct {
	Diggest    func() hash.Hash
	SaltSize   int
	KeyLen     int
	Iterations int
}

// HashResult model
type HashResult struct {
	CipherText string
	Salt       string
}

// NewPassword constructor
func NewPassword(diggest func() hash.Hash, saltSize int, keyLen int, iter int) *Password {
	return &Password{
		Diggest:    diggest,
		SaltSize:   saltSize,
		KeyLen:     keyLen,
		Iterations: iter,
	}
}

func (p *Password) genSalt() string {
	saltBytes := make([]byte, p.SaltSize)
	rand.Seed(time.Now().UnixNano())
	rand.Read(saltBytes)
	return base64.StdEncoding.EncodeToString(saltBytes)
}

// HashPassword method
func (p *Password) HashPassword(password string) HashResult {
	saltString := p.genSalt()
	salt := bytes.NewBufferString(saltString).Bytes()
	df := pbkdf2.Key([]byte(password), salt, p.Iterations, p.KeyLen, p.Diggest)
	cipherText := base64.StdEncoding.EncodeToString(df)
	return HashResult{CipherText: cipherText, Salt: saltString}
}

// VerifyPassword method
func (p *Password) VerifyPassword(password, cipherText, salt string) bool {
	saltBytes := bytes.NewBufferString(salt).Bytes()
	df := pbkdf2.Key([]byte(password), saltBytes, p.Iterations, p.KeyLen, p.Diggest)
	newCipherText := base64.StdEncoding.EncodeToString(df)
	return newCipherText == cipherText
}
