package like

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *LikeController) Delete(ctx *gin.Context) {
	rsrcType := ctx.Param("type")
	rsrcId := ctx.Param("resourceid")

	token := ctx.GetHeader("Authorization")
	if token == "" {
		core.WriteResponse(ctx, core.ErrLoginNeed, nil)
		return
	}

	opUser, err := c.Service.Users().Verify(ctx, token, &v1.VerifyOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	like := v1.Like{
		ItemType: rsrcType,
		ItemID:   rsrcId,
		UserName: opUser.UserName,
	}

	if err := c.Service.Likes().Delete(ctx, &like, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
