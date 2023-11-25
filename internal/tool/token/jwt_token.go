package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateToken token token
func GenerateToken(ttl time.Duration, payload any, secretJWTKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": payload,
		//"iss": "favaa-mitra",
	})

	signedString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("error when generate token: %s", err)
	}

	return signedString, nil

}

func getKeyFunc(secretJWTKey string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretJWTKey), nil
	}
}

// ValidateToken parse token token
func ValidateToken(tokenString string, secretJWTKey string) (jwt.MapClaims, error) {
	parse := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))

	token, err := parse.Parse(tokenString, getKeyFunc(secretJWTKey))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}

	// token expired
	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expirationTime.Unix() < time.Now().Unix() {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}

func GenerateRefreshToken(ttl time.Duration, payload any, secretJWTKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(ttl).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"sub": payload,
	})

	signedString, err := token.SignedString([]byte(secretJWTKey))
	if err != nil {
		return "", fmt.Errorf("error when generate token: %s", err)
	}

	return signedString, nil

}
