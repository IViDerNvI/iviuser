package user

import (
	"github.com/ividernvi/iviuser/internal/apiserver/service"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

type UserController struct {
	Srv service.Service
}

func NewUserController(s store.Store) *UserController {
	return &UserController{Srv: service.NewService(s)}
}
