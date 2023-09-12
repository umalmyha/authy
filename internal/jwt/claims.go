package jwt

import "github.com/golang-jwt/jwt/v5"

type Claim struct {
	Key   string
	Value any
}

type Claims struct {
	claims jwt.MapClaims
}

func NewClaims(claims ...Claim) *Claims {
	c := &Claims{claims: make(jwt.MapClaims, len(claims))}
	for _, claim := range claims {
		c.Set(claim.Key, claim.Value)
	}
	return c
}

func (c *Claims) Set(k string, v any) {
	c.claims[k] = v
}

func (c *Claims) Get(k string) (any, bool) {
	v, ok := c.claims[k]
	return v, ok
}
