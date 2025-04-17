package problem

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type ProblemController struct {
	Service service.Service
}

func NewProblemController(store store.Store) *ProblemController {
	return &ProblemController{Service: service.NewService(store)}
}
