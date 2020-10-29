package main

import (
	"net/http"
	"os"
	"time"

	"github.com/practice-userlogin/golang-userlogin-practice/middlewares"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type parent struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Children  []child `json:"children"`
}

type child struct {
	Lastname string `json:"lastname"`
}

var parents []parent
var child1 []child

func getUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"parents": parents})
}

func postUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": c.MustGet("payload").(string)})
}

type userClaim struct {
	jwt.StandardClaims
	Firstname string `json:"firstname"`
}

func loginHandler(c *gin.Context) {
	//expiration time initialization; this means expires in 1 hour
	expirationTime := time.Duration(1) * time.Hour

	//making claims which basically is the payload
	claims := userClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "budi",
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
		Firstname: "Bob",
	}

	//creating the token, specifying the signing method
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	//creating secret key
	var jwtsecretkey = []byte("inijwtsecretdsadasf")

	//signing the token with secret key
	signedToken, err := token.SignedString(jwtsecretkey)

	//error catcher
	if err != nil {
		result := gin.H{
			"message": "Wrong ID/Password",
			"error":   err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
	}

	//response
	c.JSON(http.StatusOK, signedToken)
}

func main() {
	//seed data for testing nested map (objects)
	child1 = append(child1, child{Lastname: "John"})
	child1 = append(child1, child{Lastname: "Mike"})
	parents = append(parents, parent{Firstname: "Bob", Lastname: "Marley", Children: child1})

	//initializing server
	server := gin.New()

	//recovery and logger middleware
	server.Use(gin.Recovery(), gin.Logger())

	//no middleware test
	server.GET("/", loginHandler)

	//this is setting up middleware
	apiRoutes := server.Group("/api", middlewares.Authentication())
	{
		apiRoutes.POST("/user", postUserHandler)
	}

	//port initialization;
	//if env variable PORT not available, set port to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//equivalent of app.listen
	server.Run(":" + port)
}
