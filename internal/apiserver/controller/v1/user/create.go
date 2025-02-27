package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) Create(ctx *gin.Context) {
	var user v1.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	if err := c.Srv.Users().Create(ctx, &user, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
