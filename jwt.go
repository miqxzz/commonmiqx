package common

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTUtil struct {
	secretKey string
}

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func NewJWTUtil(secretKey string) *JWTUtil {
	return &JWTUtil{secretKey: secretKey}
}

func (j *JWTUtil) GenerateToken(userID int, role string) (string, error) {
	expirationTime := time.Now().Add(14 * time.Hour) // Example: 14 hours expiration
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), // Add IssuedAt
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTUtil) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GetUserIDFromToken extracts the user ID from a valid JWT token.
func (j *JWTUtil) GetUserIDFromToken(tokenString string) (int, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

// GetRoleFromToken extracts the user role from a valid JWT token.
func (j *JWTUtil) GetRoleFromToken(tokenString string) (string, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
