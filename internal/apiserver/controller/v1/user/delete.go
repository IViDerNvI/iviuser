package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) Delete(ctx *gin.Context) {
	username := ctx.Param("id")

	opUserName := ctx.MustGet("X-Operation-User-Name").(string)
	opUserStatus := ctx.MustGet("X-Operation-User-Status").(string)

	if opUserStatus != "admin" && opUserName != username {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	err := c.Srv.Users().Delete(ctx, username, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
