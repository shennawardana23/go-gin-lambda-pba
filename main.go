package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/gin-gonic/gin"
)

func getHealthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Healthcheck")
}

func logRequest(c *gin.Context) {
	fmt.Println("MASUK FUNC logRequest Inside")

	request := c.Request
	fmt.Printf("Request: %s %s %s", request.Method, request.URL.Path, request.RemoteAddr)
}

func router() *gin.Engine {
	// set gin mode
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	r := gin.Default()

	// global middleware
	r.Use(gin.Recovery())

	// Log requests
	r.Use(logRequest)

	fmt.Println("MASUK FUNC logRequest")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/healthcheck", getHealthcheck)

	return r
}

func main() {
	router := router()
	env := os.Getenv("ENV")
	port := os.Getenv("MPORT")

	if env == "prod" {
		fmt.Println("running lambda mode")
		log.Fatal(gateway.ListenAndServe(port, router))
	} else {
		fmt.Printf("running on localhost%s \n", port)
		router.Run("localhost:8080")
	}
}
