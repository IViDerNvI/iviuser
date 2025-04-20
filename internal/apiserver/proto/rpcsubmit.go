package proto

import (
	"context"
	"strconv"

	pb "github.com/ividernvi/iviuser/internal/apiserver/proto/submit"
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type SubmitRPCServer struct {
	pb.UnimplementedSubmitRPCServiceServer
	Srv service.Service
}

func NewSubmitRPCServer(store store.Store) *SubmitRPCServer {
	return &SubmitRPCServer{
		Srv: service.NewService(store),
	}
}

func (s *SubmitRPCServer) UpdateSubmit(ctx context.Context, in *pb.UpdateSubmitRequest) (out *pb.UpdateSubmitResponse, err error) {

	insId, err := strconv.Atoi(in.SubmitInstanceId)
	if err != nil {
		return nil, err
	}

	submit, err := s.Srv.Submits().Get(ctx, uint(insId), nil)
	if err != nil {
		return nil, err
	}

	submit.Override(in.ToSubmit())
	err = s.Srv.Submits().Update(ctx, submit, nil)
	if err != nil {
		return nil, err
	}

	out = &pb.UpdateSubmitResponse{
		Code:    200,
		Message: "success",
	}
	return
}
