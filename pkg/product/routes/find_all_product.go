package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	query := ctx.Query("categories")
	fmt.Println(query)
	res, err := c.FindAllProduct(context.Background(), &pb.FindAllProductRequest{
		Categories: query,
	})
	// fmt.Println("This is the response data", res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Status == 409 {
		ctx.JSON(http.StatusConflict, res)
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}
