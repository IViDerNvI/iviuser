package like

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type LikeController struct {
	Service service.Service
}

func NewLikeController(store store.Store) *LikeController {
	return &LikeController{Service: service.NewService(store)}
}
