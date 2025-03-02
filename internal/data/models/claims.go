package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Key string `json:"key"`
    UserID uint64   `json:"user_id"`
    jwt.RegisteredClaims 
}