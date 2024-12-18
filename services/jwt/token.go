package jwt

import (
	"RPJ_Overseas_Exim/go_mod_home/utils"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(secretKey []byte, email string) string {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
        "email": email,
    })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return ""
    }

    return tokenString
}

func VerifyToken(secretKey []byte, tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err!=nil {
        return "", err
    }

    if !token.Valid {
        return "", &utils.HTTPException{Message: "Token is invalid"}
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", &utils.HTTPException{Message: "Token is invalid"}
    }

    email, ok := claims["email"].(string)
    if !ok{
        return "", &utils.HTTPException{Message: "Token is invalid"}
    }

    return email, nil
}

