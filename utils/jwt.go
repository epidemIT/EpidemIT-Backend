package utils

import "github.com/dgrijalva/jwt-go/v4"

var SecretKey = "$2a$12$GzWUZf5a1iR3wJ0Iq4i3TeGB8vAWhrQ.lS4vf5zDp39aQrkD3GJva"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return webToken, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, err
	}

	return claims, nil
}