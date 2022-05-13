package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func LoginData(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody {}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.RequestLogin{
		Email: b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}