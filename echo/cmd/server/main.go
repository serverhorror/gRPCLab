package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/serverhorror/go-playground/gRPCLab/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	certFile = flag.String("certFile", "", "Server Certificate to use")
	keyFile  = flag.String("keyFile", "", "Server Key to use")
)

type EchoServer struct {
	keyFile  string
	certFile string
}

func NewEchoServer(keyFile string, certFile string) (*EchoServer, error) {
	e := &EchoServer{
		keyFile:  keyFile,
		certFile: certFile,
	}
	return e, nil
}

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
	log.SetPrefix("gRPCLab::echo::server ")
}

func main() {

	flag.Parse()

	srv, err := NewEchoServer(*certFile, *keyFile)
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", "[::1]:8000")
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(srv.certFile, srv.keyFile)
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServer(grpcServer, srv)
	grpcServer.Serve(listener)
}
