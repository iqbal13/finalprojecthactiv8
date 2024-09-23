package helpers

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	salt := 8
	pass := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(pass, salt)

	return string(hash)
}

func Compare(hashedPassword, password []byte) bool {
	hash, pass := []byte(hashedPassword), []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}

func GenerateToken(id string, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	signedToken, _ := parseToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("harap login terlebih dahulu")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

type ResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
