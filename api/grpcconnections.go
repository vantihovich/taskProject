package grpcconn

import (
	"context"
	"fmt"
	"log"
	"net"

	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

var Cli ps.GetCredsClient

func GrpcCliConn() (c ps.GetCredsClient) {
	fmt.Println("Старт gRPC клиента2")
	conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	Cli = ps.NewGetCredsClient(conn)
	return Cli
}

type TokenGeneratorServiceServer struct {
}

func (s TokenGeneratorServiceServer) GenerateToken(c context.Context, req *ps.Request) (*ps.Response, error) {
	c = context.Background()
	var err error

	//connect_to_DB(){}

	//generate token(if check = 1)(token, expires_at){}

	response := new(ps.Response)
	response.Token, response.ExpiresAt = req.Email, req.Password

	//check (connect_to_DB, user req.Email, password req.Password)( combExists bool){}

	fmt.Println("Параметры генерации токена на сервере:", response.Token, response.ExpiresAt)

	return response, err
}

func GrpcServConn() {
	fmt.Println("Старт сервера2")
	server := grpc.NewServer()
	instance := new(TokenGeneratorServiceServer)
	ps.RegisterGetCredsServer(server, instance)

	listener, err := net.Listen("tcp", ":3500")
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}

}
