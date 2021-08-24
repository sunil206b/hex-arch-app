package grpc

import (
	"github.com/sunil206b/hex-arc-app/internal/adapters/framework/left/grpc/pb"
	"github.com/sunil206b/hex-arc-app/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (ad *Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}
	arithmeticServiceServe := ad
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServe)
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
