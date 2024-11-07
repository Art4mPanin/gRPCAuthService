package jwt

import (
	"fmt"
	myerrors "github.com/Art4mPanin/gRPCAuthService/internal/errors"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func GetToken(authHeader string) (*jwt.Token, error) {
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token from get: %w", err)
	}
	return token, nil
}
func ValidateToken(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token claims")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if float64(time.Now().Unix()) > exp {
			return 0, fmt.Errorf("token expired")
		}
	} else {
		return 0, fmt.Errorf("invalid token claims")
	}
	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token subject")
	}
	sub := int(subFloat)

	return sub, nil
}
func CreateJWT(userID uint, expireTime int64, superuser bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["exp"] = expireTime
	claims["superuser"] = superuser
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
func CreateToken(userid int, superuser bool) (accessToken, refreshToken string, err error) {
	accessTokenExpireTime := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err = CreateJWT(uint(userid), accessTokenExpireTime, superuser)
	if err != nil {
		log.Printf("Failed to create access token: %s", err)
		return "", "", myerrors.SignTokenError{}
	}

	refreshTokenExpireTime := time.Now().Add(time.Hour * 24 * 30).Unix()
	refreshToken, err = CreateJWT(uint(userid), refreshTokenExpireTime, superuser)
	if err != nil {
		log.Printf("Failed to create refresh token: %s", err)
		return "", "", myerrors.SignTokenError{}
	}

	return accessToken, refreshToken, nil
}
