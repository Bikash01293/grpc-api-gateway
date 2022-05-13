package auth

import (
	"api-gateway/pkg/auth/routes"
	"github.com/gin-gonic/gin"
)


func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.RegisterData(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.LoginData(ctx, svc.Client)
}
