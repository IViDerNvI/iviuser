package submit

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type SubmitController struct {
	Service service.Service
}

func NewSubmitController(store store.Store) *SubmitController {
	return &SubmitController{
		Service: service.NewService(store),
	}
}
