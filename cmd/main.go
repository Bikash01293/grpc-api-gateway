package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"log"
	"api-gateway/pkg/product"
	"api-gateway/pkg/cart"
	"api-gateway/pkg/order"
	"github.com/gin-gonic/gin"

)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	r := gin.Default()

	k := CORSMiddleware()

	r.Use(k)

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	cart.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}