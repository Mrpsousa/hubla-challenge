package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

var cfg *Conf

type Conf struct {
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth    *jwtauth.JWTAuth
}

func NewConfig() *Conf {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	expires, _ := strconv.Atoi((os.Getenv("JWT_EXPIRESIN")))

	return &Conf{
		JWTSecret:    os.Getenv("JWT_SECRET"),
		TokenAuth:    jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil),
		JWTExpiresIn: expires,
	}
}
