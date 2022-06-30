package product

import (
	"api-gateway/pkg/auth"

	"api-gateway/pkg/config"

	"api-gateway/pkg/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient){
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.GET("/all", svc.FindAllProduct)
	routes.PUT("/update/:id", svc.UpdateProduct)
	routes.DELETE("delete/:id", svc.DeleteProduct)
	
}

func(svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func(svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func(svc *ServiceClient) FindAllProduct(ctx *gin.Context) {
	routes.FindAllProduct(ctx, svc.Client)
}

func(svc *ServiceClient) UpdateProduct(ctx *gin.Context) {
	routes.UpdateProduct(ctx, svc.Client)
}

func(svc *ServiceClient) DeleteProduct(ctx *gin.Context) {
	routes.DeleteProduct(ctx, svc.Client)
}