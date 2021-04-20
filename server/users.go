package main

import (
	//"encoding/json"
	"context"
	"fmt"
	"log"

	"net"

	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

type TokenGeneratorServiceServer struct {
}

func (s *TokenGeneratorServiceServer) GenerateToken(c context.Context, req *ps.Request) (*ps.Response, error) {
	c = context.Background()
	var err error
	response := new(ps.Response)

	response.Token, response.ExpiresAt = req.Email, req.Password

	fmt.Println("Параметры генерации токена на сервере:", response.Token, response.ExpiresAt)

	return response, err
}

func main() {
	fmt.Println("Старт сервера")
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
