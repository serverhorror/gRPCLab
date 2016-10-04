package main

import (
	"log"

	pb "github.com/serverhorror/go-playground/gRPCLab/echo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// type EchoClient struct{}

// func Echo(ctx context.Context, in *pb.Request, opts ...grpc.CallOption) (*pb.Response, error) {
// 	return nil, nil
// }

func main() {
	// client := EchoClient{}

	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithInsecure(),
		grpc.WithUserAgent("gRPCLab echo"),
	)
	conn, err := grpc.Dial("[::1]:8000", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	request := new(pb.Request)
	request.Message = "ping"

	resp, err := client.Echo(context.Background(), request)
	if err != nil {
		panic(err)
	}

	log.Printf("Resp: %v", resp)
}
