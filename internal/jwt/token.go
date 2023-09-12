package jwt

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	token  *jwt.Token
	claims *Claims
}

func (t *Token) Signed() string {
	return t.token.Raw
}

func (t *Token) Claims() *Claims {
	return t.claims
}

func (t *Token) Alg() string {
	return t.token.Method.Alg()
}
