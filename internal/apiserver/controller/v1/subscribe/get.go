package subscribe

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubscribeController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	sub, err := c.Service.Subscribes().Get(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, sub)
}
