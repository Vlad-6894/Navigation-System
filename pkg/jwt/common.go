package pkg_jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokensPair struct {
	AccessToken  string
	RefreshToken string
}

type AccessTokenClaims struct {
	UserID   string
	Role     string
	SchoolID int
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID string
	jwt.RegisteredClaims
}

var (
	envJwt    = "JWT_SECRET"
	accessTTL = 30 * time.Minute
)

func GenerateTokensPare(
	userID string,
	role string,
	schoolID int,
) (TokensPair, error, AccessTokenClaims) {
	secret := os.Getenv(envJwt)
	if secret == "" {
		return TokensPair{}, fmt.Errorf("secret key is empty!"), AccessTokenClaims{}
	}

	accessClaims := AccessTokenClaims{
		UserID:   userID,
		Role:     role,
		SchoolID: schoolID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret)
	if err != nil {
		return TokensPair{}, fmt.Errorf("fail sign access token: %w", err), AccessTokenClaims{}
	}

	refreshClaims := RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secret)
	if err != nil {
		return TokensPair{}, fmt.Errorf("fail sign refresh token: %w", err), AccessTokenClaims{}
	}

	tokenPair := TokensPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil, accessClaims
}

func GenerateAccessToken(
	userID string,
	role string,
	schoolID int,
) (string, error) {
	secret := os.Getenv(envJwt)
	if secret == "" {
		return "", fmt.Errorf("secret key is empty!")
	}

	accessClaims := AccessTokenClaims{
		UserID:   userID,
		Role:     role,
		SchoolID: schoolID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("fail sign access token: %w", err)
	}

	return accessToken, nil
}
