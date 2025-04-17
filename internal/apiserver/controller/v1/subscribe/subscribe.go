package subscribe

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type SubscribeController struct {
	Service service.Service
}

func NewSubscribeController(store store.Store) *SubscribeController {
	return &SubscribeController{Service: service.NewService(store)}
}
