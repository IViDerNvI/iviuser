package solution

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type SolutionController struct {
	service service.Service
}

func NewSolutionController(store store.Store) *SolutionController {
	return &SolutionController{
		service: service.NewService(store),
	}
}
