package grpcconn

import (
	"context"
	"fmt"
	"log"
	"net"

	in "github.com/vantihovich/taskProject/internal"
	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

var Cli ps.GetCredsClient

func GrpcCliConn() (c ps.GetCredsClient) {
	conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	Cli = ps.NewGetCredsClient(conn)
	return Cli
}

func GrpcServConn() {

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

type TokenGeneratorServiceServer struct {
}

func (s TokenGeneratorServiceServer) GenerateToken(c context.Context, req *ps.Request) (*ps.Response, error) {
	c = context.Background()
	var err error
	response := new(ps.Response)

	fmt.Println("Параметры принятые сервером:", req.Email, req.Password)
	t := in.Check(req.Email, req.Password)

	fmt.Println(" креды найдены или нет", t)

	//generate token(if check = 1)(token, expires_at){}

	response.Token, response.ExpiresAt = req.Email, req.Password

	fmt.Println("Параметры генерации токена на сервере:", response.Token, response.ExpiresAt)

	return response, err
}
