package apiserver

import (
	"log"
	"net"

	"github.com/ividernvi/iviuser/internal/apiserver/config"
	"github.com/ividernvi/iviuser/internal/apiserver/proto"
	pb "github.com/ividernvi/iviuser/internal/apiserver/proto/user"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	proto.UserRPCServer
}

func NewGRPCServer(store store.Store) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(server, &GRPCServer{
		*proto.NewUserRPCServer(store),
	})
	reflection.Register(server)
	return server
}

func RunGRPCServer(cfg *config.Config) {
	lis, err := net.Listen("tcp", cfg.Options.RPCServeOptions.Port())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := NewGRPCServer(store.Factory())
	log.Printf("gRPC server listening on %s", cfg.Options.RPCServeOptions.Port())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
