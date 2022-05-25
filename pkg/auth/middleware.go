package auth

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct{
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig{
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token)< 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK{
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Set("userId",res.UserId)
	ctx.Next()
}