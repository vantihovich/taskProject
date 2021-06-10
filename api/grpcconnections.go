package grpcconn

import (
	"context"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	ps "github.com/vantihovich/taskProject/proto"
	"google.golang.org/grpc"
)

var Cli ps.GetCredsClient

func GrpcCliConn() (c ps.GetCredsClient) {
	conn, err := grpc.Dial("127.0.0.1:3500", grpc.WithInsecure())
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Panic("Unable to create grpcCliConn:")
	}
	Cli = ps.NewGetCredsClient(conn)
	return Cli
}

func GrpcServConn() {
	server := grpc.NewServer()
	instance := new(ps.UnimplementedGetCredsServer)
	ps.RegisterGetCredsServer(server, instance)

	listener, err := net.Listen("tcp", ":3500")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Panic("Unable to create grpc listener:")
	}

	if err = server.Serve(listener); err != nil {
		log.WithFields(log.Fields{"error": err}).Panic("Unable to start server")
	}
	fmt.Println("Struct:", instance)
}

type UnimplementedGetCredsServer struct {
}

//TODO move GenerateToken func to another package
func (u UnimplementedGetCredsServer) GenerateToken(c context.Context, req *ps.Request) (*ps.Response, error) {
	c = context.Background()
	var err error
	response := new(ps.Response)

	log.WithFields(log.Fields{"Username": req.Email, "password": req.Password}).Info("Parameters received by server")

	response.Token, response.ExpiresAt = req.Email, req.Password

	log.WithFields(log.Fields{"token": response.Token, "expires_at": response.ExpiresAt}).Info("Parameters for generating token")

	return response, err
}

func (u UnimplementedGetCredsServer) mustEmbedUnimplementedGetCredsServer() {}
