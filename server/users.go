package main

import (
	"encoding/json"

	"log"
	"net/http"

	"google.golang.org/grpc"
	 ps "C:/Users/v.antsikhovich/Projects/TaskProject/proto"
)

type TokenGeneratorServiceServer struct {
}

func (s *TokenGeneratorServiceServer) GenerateToken(ctx context.Context,
    req *ps.Request) (*ps.Response, error){
    	var err error
		response:= new(ps.Response)
		
		response.token, response.expires_at = req.email, req.password
		
		fmt.Println("Параметры генерации токена на сервере:",response.token, response.expires_at  )
		
		return response,err
    }

func main() {
    server := grpc.NewServer()

    instance := new(TokenGeneratorServiceServer)

    ps.RegisterTokenGeneratorServiceServer(server, instance)

    listener, err := net.Listen("tcp", ":3500")
    if err != nil {
        log.Fatal("Unable to create grpc listener:", err)
    }

    if err = server.Serve(listener); err != nil {
        log.Fatal("Unable to start server:", err)
    }
}