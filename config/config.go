package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct{
	DbUrl string
	Port string
	JWTSecert string
}
func LoadConfig() *config{
	godotenv.Load()
	dbUrl:=os.Getenv("DatabaseUrl")
	Port:=os.Getenv("PORT")
	Jwt:=os.Getenv("JWTSecert")

	return &config{
		DbUrl:dbUrl,
		Port: Port,
		JWTSecert: Jwt,
	}

}