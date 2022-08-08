package cart

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/auth"
	"api-gateway/pkg/cart/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient {
		Client: InitServiceClient(c),
	}

	routes := r.Group("/cart")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateCart)

}

func (svc *ServiceClient) CreateCart(ctx *gin.Context) {
	routes.CreateCart(ctx, svc.Client)
}