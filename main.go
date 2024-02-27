package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func getHealthcheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}

func router() *gin.Engine {
	// set gin mode
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	r := gin.Default()

	// global middleware
	r.Use(gin.Recovery())

	// r.HandleMethodNotAllowed = true // healtcheck

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
