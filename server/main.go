package main

import (
	"context"
	"log"
	"net"

	add "github.com/souvikhaldar/grpc-tutorial/add"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) Adder(ctx context.Context, req *add.AdderRequest) (*add.AdderResponse, error) {
	var resp add.AdderResponse
	resp.AddResult = req.FirstNum + req.SecondNum
	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8192")
	if err != nil {
		log.Fatal("Failed to listen to TCP connections: ", err)
	}
	server := grpc.NewServer()
	add.RegisterAddServer(server, &Server{})
	server.Serve(lis)
}
