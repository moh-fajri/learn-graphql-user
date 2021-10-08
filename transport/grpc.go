package transport

import (
	"github.com/moh-fajri/learn-garphql-user/action"
	pb "github.com/moh-fajri/learn-garphql-user/entity/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func (svc Service) RunGRCP() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Connect GRPC to port " + os.Getenv("PORT"))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, action.NewHandler())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
