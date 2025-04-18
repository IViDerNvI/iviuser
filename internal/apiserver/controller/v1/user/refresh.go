package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
	"github.com/ividernvi/iviuser/pkg/util/jwtutil"
)

func (c *UserController) Refresh(ctx *gin.Context) {

	if c.Srv.Users().Logout(ctx, ctx.GetHeader("Authorization"), nil) != nil {
		core.WriteResponse(ctx, core.ErrTokenInvalid, nil)
		return
	}

	opUserNameRaw, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserNameRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserName, ok := opUserNameRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserStatusRaw, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatusRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserStatus, ok := opUserStatusRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserInstanceIdRaw, ok := ctx.Get("X-Operation-User-InstanceID")
	if !ok || opUserInstanceIdRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserInstanceId, ok := opUserInstanceIdRaw.(uint)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	jwt, err := jwtutil.CreateJWT(&v1.User{
		UserName: opUserName,
		Status:   opUserStatus,
		ObjMeta: v1.ObjMeta{
			InstanceID: opUserInstanceId,
		},
	})
	if err != nil {
		core.WriteResponse(ctx, core.ErrTokenCreateFailed, nil)
		return
	}

	ctx.Header("Authorization", jwt)
	core.WriteResponse(ctx, nil, jwt)
}
