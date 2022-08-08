package routes

import (
	"api-gateway/pkg/cart/pb"
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type CreateCartRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateCart(ctx *gin.Context, c pb.CartServiceClient) {
	b := CreateCartRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	fmt.Println("This is the userId", reflect.TypeOf(userId))
	res, err := c.CreateCart(context.Background(), &pb.CreateCartRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
