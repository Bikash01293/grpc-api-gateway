package routes

import (
	"api-gateway/pkg/order/pb"
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	b := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	fmt.Println("This is the userId", reflect.TypeOf(userId))
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
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
