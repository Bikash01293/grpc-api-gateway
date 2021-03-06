package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductRequestBody struct {
	Title  string `json:"title"`
	Desc string  `json:"desc"`
	Img string  `json:"img"`
	Categories []string `json:"categories"`
	Size  string `json:"size"`
	Color  string `json:"color"`
	Price  int64 `json:"price"`
	Stock  int64 `json:"stock"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	b := CreateProductRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(b)
	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Title:  b.Title,
		Desc: b.Desc,
		Img: b.Img,
		Categories: b.Categories,
		Size: b.Size,
		Color: b.Color,
		Price: b.Price,
		Stock: b.Stock,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}
