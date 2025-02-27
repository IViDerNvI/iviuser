package proto

import (
	"context"

	pb "github.com/ividernvi/iviuser/internal/apiserver/proto/user"
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type UserRPCServer struct {
	pb.UnimplementedUserInfoServiceServer
	Srv service.Service
}

func NewUserRPCServer(store store.Store) *UserRPCServer {
	return &UserRPCServer{
		Srv: service.NewService(store),
	}
}

func (s *UserRPCServer) Create(ctx context.Context, in *pb.UserInfo) (*pb.ErrorResponse, error) {
	err := s.Srv.Users().Create(ctx, pb.ToUser(in), nil)
	return pb.ToErrorResponse(err), err
}

func (s *UserRPCServer) Get(ctx context.Context, in *pb.GetUserRequest) (*pb.UserInfo, error) {
	u, err := s.Srv.Users().Get(ctx, in.Username, nil)
	if err != nil {
		return nil, err
	}
	return pb.ToUserInfo(u), nil
}

func (s *UserRPCServer) Update(ctx context.Context, in *pb.UserInfo) (*pb.ErrorResponse, error) {
	err := s.Srv.Users().Update(ctx, pb.ToUser(in), nil)
	return pb.ToErrorResponse(err), err
}

func (s *UserRPCServer) Delete(ctx context.Context, in *pb.GetUserRequest) (*pb.ErrorResponse, error) {
	err := s.Srv.Users().Delete(ctx, in.Username, nil)
	return pb.ToErrorResponse(err), err
}
