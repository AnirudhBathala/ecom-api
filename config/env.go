package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database_URI string
	JWTSecret string
	JWTExpirationInSeconds int64
}

var Configs = InitConfig()

func InitConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file",err)
	}

	expiration,err:=strconv.ParseInt( os.Getenv("JWTEXPIRATIONINSECONDS"),10,64)
	if err!=nil {
		// fallback incase if strconv fails
		expiration = 3600 * 24 * 7
	}

	return Config{
		Database_URI: os.Getenv("DATABASE_URI"),
		JWTSecret: os.Getenv("JWTSECRET"),
		JWTExpirationInSeconds: expiration ,
	}
}
