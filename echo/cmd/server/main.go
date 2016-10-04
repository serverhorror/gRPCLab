package main

import (
	"log"
	"net"

	pb "github.com/serverhorror/go-playground/gRPCLab/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoServer struct{}

func (e EchoServer) Echo(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received a message(%T): %v", in, in)
	defer log.Print("Done")

	resp := new(pb.Response)
	resp.Message = in.Message
	if in.Message == "ping" {
		resp.Message = "pong"
	}
	return resp, nil
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.SetPrefix("gRPCLab::echo ")
}

func main() {
	srv := EchoServer{}
	listener, err := net.Listen("tcp", "[::1]:8000")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServer(grpcServer, srv)
	grpcServer.Serve(listener)
}
