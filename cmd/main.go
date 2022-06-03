package main

import (
	"api-gateway/pkg/auth"
	"api-gateway/pkg/config"
	"log"

	"api-gateway/pkg/order"
	"api-gateway/pkg/product"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to load config", err)
	}

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	r.Use(cors.Default())

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
