package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secretKey []byte, username string) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
        "username": username,
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return ""
    }

    return tokenString
}

func VerifyToken(secretKey []byte, tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })

    if err!=nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("Token is invalid")
    }

    return token, nil
}
