package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load config", err)
	}
	
	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
