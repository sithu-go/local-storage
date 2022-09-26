package filegrpc

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Conn = &grpc.ClientConn{}
)

func Connectgrpc() {
	c, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("Cannot connect grpc", err)
		return
	}
	fmt.Println("GRPC Connected")
	Conn = c
}
