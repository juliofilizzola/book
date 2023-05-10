package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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

func GetUserId(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, err := jwt.Parse(tokenString, getKey)
	if err != nil {
		return 0, err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		permissionSting := fmt.Sprintf("%.0f", permission["userId"])
		userId, err := strconv.ParseUint(permissionSting, 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}

	return 0, errors.New("token invalid")
}
