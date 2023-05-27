package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

func GenerateToken(userId string) (string, error) {
	fmt.Println(userId)
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["userId"] = userId

	fmt.Println(permission)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	s, _ := token.SignedString([]byte(config.SecretKey))
	return s, nil
}

func ValidToken(r *http.Request) error {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getKey)

	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalid")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	formatToken := strings.Split(token, " ")
	if len(formatToken) == 2 {
		return formatToken[1]
	}
	return ""
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("method incorrect")
	}

	return []byte(config.SecretKey), nil
}

func GetUserId(r *http.Request) (string, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return "", err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := fmt.Sprintf("%v", permission["userId"])
		return userId, nil
	}

	return "", errors.New("token invalid")
}
