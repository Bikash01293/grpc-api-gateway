package product

import (
	"api-gateway/pkg/config"
	"api-gateway/pkg/product/pb"
	"log"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("unable to generate grpc connection", err)
	}

	return pb.NewProductServiceClient(cc)
}