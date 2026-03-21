package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	DbUrl string
	Port string
	JWTSecert string
}
func LoadConfig() *Config{
	godotenv.Load()
	dbUrl:=os.Getenv("DATABASE_URL")
	Port:=os.Getenv("PORT")
	Jwt:=os.Getenv("JWT_SECERT")

	return &Config{
		DbUrl:dbUrl,
		Port: Port,
		JWTSecert: Jwt,
	}

}