package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const ExpiresIn = time.Hour * 1

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn ...time.Duration) (string, error) {
	expiration := time.Hour
	if len(expiresIn) > 0 && expiresIn[0] > 0 {
		expiration = expiresIn[0]
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "ACCTJ",
		IssuedAt:  &jwt.NumericDate{time.Now()},
		ExpiresAt: &jwt.NumericDate{time.Now().Add(expiration)},
		Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return fmt.Sprintf("Error signing token with secret : %s", err), err
	}

	return signed, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		// This is a security check to ensure the token was signed with the expected algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// The parser will use this key to verify the token's signature
		return []byte(tokenSecret), nil
	})

	// Handle parsing errors
	if err != nil {
		return uuid.Nil, fmt.Errorf("error parsing token: %w", err)
	}

	//Check if token is valid and claims are properly parsed
	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	user, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, err
	}

	return user, nil
}
