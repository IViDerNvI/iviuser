package like

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"

	v1 "github.com/ividernvi/iviuser/model/v1"
)

func (c *LikeController) Get(ctx *gin.Context) {
	rsrcId := ctx.Param("resourceid")
	rsrcType := ctx.Param("type")

	like := &v1.Like{
		ItemType: rsrcType,
		ItemID:   rsrcId,
	}

	result, err := c.Service.Likes().Get(ctx, like, &v1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, result)
}
