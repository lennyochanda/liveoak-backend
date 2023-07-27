package tokenutil

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/lennyochanda/LiveOak/types"
)

type JWTClaim struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type JWTRefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func CreateAccessToken(user *types.User, secret string, expiry int) (accessToken string, err error) {
	expiration := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &JWTClaim{
		ID:       user.ID,
		Email:    user.Email,
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, err
}

func CreateRefreshToken(user *types.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &JWTRefreshClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, err
}

func IsValid(requestToken string, secret string) (bool, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, fmt.Errorf("not a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			return false, fmt.Errorf("expired token")
		} else {
			return false, fmt.Errorf("invalid token")
		}
	}


	return false, fmt.Errorf("invalid token")
}

func ExtractID(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "error:", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "error:", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}
