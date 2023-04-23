package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"time"
)

const (
	salt       = "asklfjn2jdnalkmsd"
	signingKey = "adSj23&h#!kjWjqwnd@jnef7832N"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId  string `json:"user_id"`
	IsAdmin bool   `json:"isAdmin"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) GenerateTokenForAgent(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, generatePassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Login,
		false,
	})

	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) GenerateTokenForAdmin() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		viper.GetString("admin.login"),
		true,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, bool, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("tüzsüz signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "-", false, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok && token.Valid {
		return "-", false, errors.New("token claims'leri *tokenClaims örnegi bolup tügüldür")
	}
	return claims.UserId, claims.IsAdmin, nil
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
