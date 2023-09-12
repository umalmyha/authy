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
	Issuer            string
	Method            jwt.SigningMethod
	SigningKeyFunc    SigningKeyFunc
	ValidationKeyFunc ValidationKeyFunc
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

func (i *Issuer) Sign(mc Claims) (string, error) {
	claims := make(jwt.MapClaims, len(mc))
	for k, v := range mc {
		claims[k] = v
	}

	claims["iss"] = i.cfg.Issuer
	claims["kid"] = uuid.NewString()

	token := jwt.NewWithClaims(i.cfg.Method, claims)

	kid := i.cfg.SigningKeyFunc()
	kp, ok := i.keys[kid]
	if !ok {
		return "", fmt.Errorf("failed to sign JWT: key pair with id %s is missing", kid)
	}

	signed, err := token.SignedString(kp.PrivateKey)
	if err != nil {
		return "", err
	}

	return signed, nil
}

func (i *Issuer) Parse(raw string) error {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(
		raw,
		claims,
		func(t *jwt.Token) (any, error) {
			return i.cfg.ValidationKeyFunc()
		}
		i.cfg.ValidationKeyFunc,
		jwt.WithIssuer(i.cfg.Issuer),
		jwt.WithValidMethods([]string{i.cfg.Method.Alg()}),
		jwt.WithIssuedAt(),
	)
	if err != nil {
		return err
	}
}
