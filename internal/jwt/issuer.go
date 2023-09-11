package jwt

import (
	"crypto"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SigningKeyFunc func() string

type ValidationKeyFunc = func(*Token) (any, error)

type KeyPair struct {
	ID         string
	PrivateKey crypto.PrivateKey
	PublicKey  crypto.PublicKey
}

type IssuerConfig struct {
	issuer      string
	method      jwt.SigningMethod
	signKeyFunc SigningKeyFunc
}

type Claims struct {
	jwt.RegisteredClaims
}

type Issuer struct {
	keys map[string]KeyPair
	cfg  IssuerConfig
}

func NewIssuer(keys []KeyPair, cfg IssuerConfig) *Issuer {
	km := make(map[string]KeyPair, len(keys))
	for _, key := range keys {
		km[key.ID] = key
	}

	return &Issuer{
		keys: km,
		cfg:  cfg,
	}
}

func (i *Issuer) Sign(c MapClaims) error {
	rc := MapClaims{
		"iss": i.cfg.issuer,
		"kid": uuid.NewString(),
	}

	token := jwt.NewWithClaims(i.cfg.method, rc)

	kid := i.cfg.signKeyFunc()
	kp, ok := i.keys[kid]
	if !ok {
		return fmt.Errorf("failed to sign JWT: key pair with id %s is missing", kid)
	}

	token.SignedString()

}
