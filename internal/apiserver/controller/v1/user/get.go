package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) Get(ctx *gin.Context) {
	username := ctx.Param("id")
	if username == "" {
		core.WriteResponse(ctx, errors.New("username need"), nil)
		return
	}

	user, err := c.Srv.Users().Get(ctx, username, nil)

	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, user)
}
