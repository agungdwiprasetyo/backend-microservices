package rsa

import (
	"crypto/rsa"
	_ "embed"

	"github.com/dgrijalva/jwt-go"
	"pkg.agungdp.dev/candi/codebase/interfaces"
)

var (
	//go:embed public.pem
	verifyBytes []byte

	//go:embed private.key
	signBytes []byte
)

type rsaKey struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// InitKey rsa
func InitKey() interfaces.RSAKey {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic("missing rsa public key file, make sure you are running `generate_rsa_key` script and put the file in configs/rsa directory")
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic("missing rsa private key file, make sure you are running `generate_rsa_key` script and put the file in configs/rsa directory")
	}

	verifyBytes, signBytes = nil, nil
	return &rsaKey{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (r rsaKey) PrivateKey() *rsa.PrivateKey {
	return r.privateKey
}

func (r rsaKey) PublicKey() *rsa.PublicKey {
	return r.publicKey
}
