package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductRequestBody struct {
	Title  string `json:"title"`
	Desc string  `json:"desc"`
	Img string  `json:"img"`
	Categories []string `json:"categories"`
	Size  string `json:"size"`
	Color  string `json:"color"`
	Price  int64 `json:"price"`
	Stock  int64 `json:"stock"`
}


func UpdateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	fmt.Println("the requested id is:", id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	b := UpdateProductRequestBody{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.UpdateProduct(context.Background(), &pb.UpdateProductRequest{
		Id: int64(id),
		Title:  b.Title,
		Desc: b.Desc,
		Img: b.Img,
		Categories: b.Categories,
		Size: b.Size,
		Color: b.Color,
		Price: b.Price,
		Stock: b.Stock,
	})
	fmt.Println("This is the response data", res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}


