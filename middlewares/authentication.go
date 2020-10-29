package middlewares

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Authentication middleware to validate JWT token and opens the payload
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("access_token")

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Invalid signing method of %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		//jwt verify and make payload
		if token != nil && err == nil {
			c.Set("payload", token)
		} else {
			result := gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
	}
}
