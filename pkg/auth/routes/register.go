package routes

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type RegisterRequestBody struct {
	Name    string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Cpassword string `json:"cpassword"`
}

func RegisterData(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Name: b.Name,
		Email:    b.Email,
		Password: b.Password,
		Cpassword: b.Cpassword,
	})
	if err != nil {
		fmt.Println("this is the error")
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// fmt.Println(res)
	ctx.JSON(http.StatusCreated, gin.H{"Data": res})

}
