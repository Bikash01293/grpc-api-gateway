package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"log"

	"github.com/gin-gonic/gin"
	"api-gateway/pkg/product"
	"api-gateway/pkg/order"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load config", err)
	}
	
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
