package subscribe

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubscribeController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.Service.Subscribes().Delete(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
