package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(userId uint64) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	return token.SignedString([]byte(config.SecretKey))
}

func ValidToken(r *http.Request) error {
	token, err := convertToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	validToken := strings.Split(token, " ")
	if len(validToken) == 2 {
		return validToken[1]
	}
	return ""
}

func validKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method invalid! %v", token.Header["alg"])
	}

	return config.SecretKey, nil

}

func getIdToken(r *http.Request) (uint64, error) {
	token, err := convertToken(r)
	if err != nil {
		return 0, err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tokenPermission = fmt.Sprintf("%.0f", permission["userId"])
		userId, err := strconv.ParseUint(tokenPermission, 10, 64)
		if err != nil {
			return 0, err
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}

func convertToken(r *http.Request) (*jwt.Token, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, validKey)
	if err != nil {
		return nil, err
	}
	return token, nil
}
