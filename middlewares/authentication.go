package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtsecretkey = []byte("inijwtsecretdsadasf")

//Authentication middleware to validate JWT token and opens the payload
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("access_token")

		//jwt verify
		/* 		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Invalid signing method")
		}
		return jwtsecretkey, nil
		}) */

		tokenString := accessToken[len("Bearer "):]
		fmt.Println("hehehey")
		fmt.Printf("tokenString: %v \n %T \n", tokenString, tokenString)

		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Signing Method")
			}
			return jwtsecretkey, nil
		})

		//claims itu istilah buat payload disini
		if parsedToken.Valid {
			result := gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		} else {
			claims, ok := parsedToken.Claims.(jwt.MapClaims)
			if ok {
				c.Set("payload", claims)
			} else {
				result := gin.H{
					"message": "Unauthorized",
					"error":   err.Error(),
				}
				c.JSON(http.StatusUnauthorized, result)
			}
		}
	}
}
