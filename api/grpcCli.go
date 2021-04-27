package grpcCli

import (
	"log"

	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

func Grpc_connect() (c ps.GetCredsClient) {
	conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c = ps.NewGetCredsClient(conn)
	return c
}
