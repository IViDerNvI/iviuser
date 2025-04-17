package user

import (
	"io"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (u *UserController) PutAvatar(ctx *gin.Context) {
	// 使用 io.ReadAll 读取请求体
	avatarBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	// 调用服务层方法，将字节流传入 MinIO
	err = u.Srv.Users().PutAvatar(ctx, ctx.Param("id"), avatarBytes, &v1.UpdateOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
