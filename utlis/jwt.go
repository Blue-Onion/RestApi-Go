package utlis

import (
	"time"

	"github.com/Blue-Onion/RestApi-Go/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJwt(userId uuid.UUID) (string,error){
	JWTSecert:=[]byte(config.LoadConfig().JWTSecert)
	claims:=jwt.MapClaims{
		"userId": userId.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
        "iat":     time.Now().Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(JWTSecert)
}