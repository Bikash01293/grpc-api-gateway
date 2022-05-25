package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	b := CreateProductRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(b)
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  b.Name,
		Stock: b.Stock,
		Price: b.Price,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
