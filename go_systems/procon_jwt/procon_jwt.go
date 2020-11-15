package procon_jwt

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// create a struct that will be encoded to a jwt
// we add jwt.standard claims as an embedded type, to provide fields like expiry time

func GenerateJWT(privkeyfile *rsa.PrivateKey, name string, alias string, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	expireTenMinutes := time.Now().Add(time.Duration(30) * time.Minute).Unix()
	token.Claims = jwt.MapClaims{
		"iss":    "vigneshsivarajah.no",
		"aud":    "localhost:3000",
		"exp":    expireTenMinutes,
		"jti":    "Unique",
		"iat":    time.Now().Unix(),
		"nbf":    2,
		"sub":    "subject",
		"scopes": "api:read,api:write",
		"name":   name,
		"alias":  alias,
		"email":  email,
		"role":   role,
	}
	tokenString, err := token.SignedString(privkeyfile)
	if err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func ValidateJWT(publickeyfile *rsa.PublicKey, jwtgo string) (bool, error) {
	token, err := jwt.Parse(jwtgo, func(token *jwt.Token) (interface{}, error) {
		return publickeyfile, nil
	})
	if err != nil {
		return false, err
	} else if token.Valid && err == nil {
		return true, nil
	}
	return false, err
}
