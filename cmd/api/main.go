package main

import (
	"log"

	"github.com/crisywini/go-products-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Load()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Server up and running on port %s", config.ServerPort)
	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatal(err)
	}
}
