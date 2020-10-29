package main

import (
	"fmt"
	"net/http"
	"os"

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

//the middleware
func authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("inside Auth")
	}
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
	server.GET("/", getUserHandler)

	//this is setting up middleware
	apiRoutes := server.Group("/api", authentication())
	{
		apiRoutes.GET("/user", getUserHandler)
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
