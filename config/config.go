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
	dbUrl:=os.Getenv("DatabaseUrl")
	Port:=os.Getenv("PORT")
	Jwt:=os.Getenv("JWTSecert")

	return &Config{
		DbUrl:dbUrl,
		Port: Port,
		JWTSecert: Jwt,
	}

}