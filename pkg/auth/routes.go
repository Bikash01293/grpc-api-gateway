package auth

import (
	"api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient{
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	// fmt.Println(svc.Client.Register(context.Background(), &pb.RequestRegister{}))

	route := r.Group("/auth")
	route.POST("/register", svc.Register)
	route.POST("/login", svc.Login)

	return svc
}
