package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
	"github.com/ividernvi/iviuser/pkg/util/jwtutil"
)

func (c *UserController) Login(ctx *gin.Context) {
	jwt, err := jwtutil.CreateJWT(&v1.User{
		UserName: ctx.MustGet("X-Operation-User-Name").(string),
		Status:   ctx.MustGet("X-Operation-User-Status").(string),
		ObjMeta: v1.ObjMeta{
			InstanceID: ctx.MustGet("X-Operation-User-InstanceID").(uint),
		},
	})
	if err != nil {
		core.WriteResponse(ctx, core.ErrTokenCreateFailed, nil)
		return
	}

	ctx.Header("Authorization", jwt)
	core.WriteResponse(ctx, nil, jwt)
}
