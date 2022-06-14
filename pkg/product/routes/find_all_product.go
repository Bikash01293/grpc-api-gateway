package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	
	res, err := c.FindAllProduct(context.Background(), &pb.FindAllProductRequest{})
	fmt.Println("This is the response data", res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
