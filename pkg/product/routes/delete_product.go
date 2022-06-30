package routes

import (
	"api-gateway/pkg/product/pb"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
	fmt.Println("the requested id is:", id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	

	res, err := c.DeleteProduct(context.Background(), &pb.DeleteProductRequest{
		Id: int64(id),
	})
	fmt.Println("This is the response data", res)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}


