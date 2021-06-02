package grpcconn

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
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
		log.WithFields(log.Fields{"error": err}).Panic("Unable to create grpc listener:")
	}

	if err = server.Serve(listener); err != nil {
		log.WithFields(log.Fields{"error": err}).Panic("Unable to start server")
	}

}

type TokenGeneratorServiceServer struct {
}

func (s TokenGeneratorServiceServer) GenerateToken(c context.Context, req *ps.Request) (*ps.Response, error) {
	c = context.Background()
	var err error
	response := new(ps.Response)

	log.WithFields(log.Fields{}).Info("Parameters received by server")
	//fmt.Println("Параметры принятые сервером:", req.Email, req.Password)

	t := in.Check(req.Email, req.Password)

	fmt.Println("If credentials are found in DB or not", t)

	//generate token(if check = 1)(token, expires_at){}

	response.Token, response.ExpiresAt = req.Email, req.Password

	log.WithFields(log.Fields{"token": response.Token, "expires_at": response.ExpiresAt}).Info("Parameters for generating token")

	return response, err
}

func (s TokenGeneratorServiceServer) mustEmbedUnimplementedGetCredsServer() {}
