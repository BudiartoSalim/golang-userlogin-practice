package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtsecretkey = []byte("inijwtsecretdsadasf")

//Authentication middleware to validate JWT token and opens the payload
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("access_token")
		tokenString := strings.Replace(accessToken, "Bearer ", "", -1)

		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Signing Method")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Invalid Signing Method")
			}
			return jwtsecretkey, nil
		})

		if err != nil {
			result := gin.H{
				"message": "Auth Fail",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)

		if !ok || !parsedToken.Valid {
			result := gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
		} else {
			c.JSON(http.StatusOK, claims)
		}
	}
}
