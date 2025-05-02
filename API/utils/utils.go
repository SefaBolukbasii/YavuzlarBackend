package utils

import (
	"API/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("gizli")

func WriteJson(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
func GenerateJwt(userId int, userName string, role string) (string, error) {
	bitisZaman := time.Now().Add(time.Hour * 1)
	claims := &models.User{
		Id:       userId,
		Username: userName,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(bitisZaman)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ValidateToken(tokenString string) (*models.User, error) {
	user := &models.User{}
	token, err := jwt.ParseWithClaims(tokenString, user, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return user, nil
}
func ExtractToken(r *http.Request) (*models.User, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, fmt.Errorf("authorization header not found")
	}
	tokenString = tokenString[len("Bearer "):]
	user, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}
	return user, nil
}
