package rsa

import (
	"crypto/rsa"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"pkg.agungdwiprasetyo.com/candi/candihelper"
	"pkg.agungdwiprasetyo.com/candi/codebase/interfaces"
)

const (
	privateKeyPath = "configs/rsa/private.key"
	publicKeyPath  = "configs/rsa/public.pem"
)

var (
	// VerifyKey rsa from config
	VerifyKey *rsa.PublicKey

	signKey *rsa.PrivateKey
)

type rsaKey struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// InitPublicKey return *rsa.PublicKey
func InitPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	verifyBytes, err := ioutil.ReadFile(os.Getenv(candihelper.WORKDIR) + publicKeyPath)
	if err != nil {
		return nil, err
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return VerifyKey, nil
}

// InitPrivateKey return *rsa.PrivateKey
func InitPrivateKey(privateKeyPath string) (*rsa.PrivateKey, error) {
	signBytes, err := ioutil.ReadFile(os.Getenv(candihelper.WORKDIR) + privateKeyPath)
	if err != nil {
		return nil, err
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, err
	}
	return signKey, nil
}

// InitKey rsa
func InitKey() interfaces.RSAKey {
	publicKey, err := InitPublicKey(publicKeyPath)
	if err != nil {
		panic("missing rsa public key file, make sure you are running `generate_rsa_key` script and put the file in configs/rsa directory")
	}
	privateKey, err := InitPrivateKey(privateKeyPath)
	if err != nil {
		panic("missing rsa private key file, make sure you are running `generate_rsa_key` script and put the file in configs/rsa directory")
	}

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
