package subscribe

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubscribeController) List(ctx *gin.Context) {
	subscribes, err := c.Service.Subscribes().List(ctx, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, subscribes)
}
