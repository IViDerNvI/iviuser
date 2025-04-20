package apiserver

import (
	"log"
	"net"

	"github.com/ividernvi/iviuser/internal/apiserver/config"
	"github.com/ividernvi/iviuser/internal/apiserver/proto"
	pbSolution "github.com/ividernvi/iviuser/internal/apiserver/proto/solution"
	pbSubmit "github.com/ividernvi/iviuser/internal/apiserver/proto/submit"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	proto.SolutionRPCServer
	proto.SubmitRPCServer
}

func NewGRPCServer(store store.Store) *grpc.Server {
	server := grpc.NewServer()

	solutionRPCServer := proto.NewSolutionRPCServer(store)
	submitRPCServer := proto.NewSubmitRPCServer(store)

	pbSolution.RegisterSolutionRPCServiceServer(server, solutionRPCServer)
	pbSubmit.RegisterSubmitRPCServiceServer(server, submitRPCServer)

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
