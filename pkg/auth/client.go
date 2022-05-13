package auth

import (
	"api-gateway/pkg/auth/pb"
	"api-gateway/pkg/config"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	// fmt.Println("this is the second error")

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	// defer conn.Close()

	return pb.NewAuthServiceClient(conn)
}
