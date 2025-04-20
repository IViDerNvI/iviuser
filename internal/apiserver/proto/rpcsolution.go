package proto

import (
	"context"

	pb "github.com/ividernvi/iviuser/internal/apiserver/proto/solution"
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
	v1 "github.com/ividernvi/iviuser/model/v1"
)

type SolutionRPCServer struct {
	pb.UnimplementedSolutionRPCServiceServer
	Srv service.Service
}

func NewSolutionRPCServer(store store.Store) *SolutionRPCServer {
	return &SolutionRPCServer{
		Srv: service.NewService(store),
	}
}

func (s *SolutionRPCServer) GetSolution(ctx context.Context, in *pb.GetSolutionRequest) (*pb.GetSolutionResponse, error) {
	mapper := map[string]string{
		"problem_id": in.ProblemId,
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Selector: selector,
	}
	listOptions.Complete()

	solutions, err := s.Srv.Solutions().List(ctx, listOptions)
	if err != nil {
		return nil, err
	}

	var solution pb.GetSolutionResponse

	for _, s := range solutions.Items {
		solution.Solutions = append(solution.Solutions, &pb.SolutionInfo{
			TestData:   s.TestData,
			TestResult: s.TestResult,
		})
	}
	return &solution, nil
}
