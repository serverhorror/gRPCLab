package main

import (
	"flag"
	"log"

	pb "github.com/serverhorror/go-playground/gRPCLab/echo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

var (
	caFile = flag.String("caFile", "", "Path to CA")
)

// type EchoClient struct{}

// func Echo(ctx context.Context, in *pb.Request, opts ...grpc.CallOption) (*pb.Response, error) {
// 	return nil, nil
// }

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.SetPrefix("gRPCLab::echo::client ")
}

func main() {
	log.Print("Started and Initialized")
	// client := EchoClient{}

	flag.Parse()

	var opts []grpc.DialOption
	var sn string

	creds, err := credentials.NewClientTLSFromFile(*caFile, sn)
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}

	opts = append(opts,
		grpc.WithTransportCredentials(creds),
		grpc.WithUserAgent("gRPCLab echo"),
	)
	conn, err := grpc.Dial("localhost:8000", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	request := new(pb.Request)
	request.Message = "ping"

	log.Printf("Sending %v", request)
	resp, err := client.Echo(context.Background(), request)
	if err != nil {
		panic(err)
	}
	log.Printf("Response: %v", resp)
	log.Print("Done Sending")

}
