package configurations

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("my_secret_key")

var claims = jwt.MapClaims{
	"authorized": true,
	"user":       "admin",
	"exp":        time.Now().Add(20 * time.Minute),
}

func NewToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (bool, error) {
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, errors.New("error signature")
		}
		return false, err
	}
	if !tkn.Valid {
		return false, errors.New("unauthorized")
	}
	return true, nil
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		var tokenArray = strings.Split(r.Header["Authorization"][0], " ")

		if len(tokenArray) != 2 {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		_, verifyError := VerifyToken(tokenArray[1])

		if verifyError != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func GenerateHash(password string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return hash, err
	}

	return hash, nil
}

func CompareHashAndPassword(password string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
