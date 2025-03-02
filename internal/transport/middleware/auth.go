package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/koyo-os/crm/internal/data/models"
)

func getClaims(tokenStr string) (*models.Claims, error) {
    // Ключ для подписи/проверки токена (должен совпадать с тем, что использовался при создании токена)
    var jwtKey = []byte(os.Getenv("JWT_SEKRET_KEY")) // Замените на ваш секретный ключ

    claims := &models.Claims{}

    // Парсим токен
    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        // Проверяем, что алгоритм подписи совпадает с ожидаемым
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}


func parseToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return os.Getenv("JWT_SEKRET_KEY"), nil
    })
}

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := w.Header().Get("Authification")

		if authHeader == "" {
			http.Error(w, "auth token empty", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := getClaims(tokenStr)
        if err != nil || claims == nil {
            http.Error(w, "token not valid", http.StatusUnauthorized) // Исправлено на 401
            return
        }

		ctx := context.WithValue(r.Context(), "claims", claims)
		r.WithContext(ctx)

		handler(w, r)
	}
}