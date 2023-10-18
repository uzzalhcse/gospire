package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/uzzalhcse/GoSpire/app/models"
	"regexp"
	"time"
)

// IsEmail Helper function to check if a string is an email address
func IsEmail(str string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(str)
}
func GenerateToken(user models.User, secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID // Subject
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
