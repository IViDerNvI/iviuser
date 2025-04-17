package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) Create(ctx *gin.Context) {
	var user v1.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if user.IsAdmin() {
		operatorStatus := ctx.MustGet("X-Operation-User-Status").(string)
		if operatorStatus != "admin" {
			core.WriteResponse(ctx, core.ErrAdminNeed, nil)
			return
		}
	}

	if err := user.Validate(); err != nil {
		core.WriteResponse(ctx, core.ErrUserInvalid, nil)
		return
	}

	if err := c.Srv.Users().Create(ctx, &user, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseCreate, nil)
		return
	}

	core.WriteResponse(ctx, nil, user)
}
