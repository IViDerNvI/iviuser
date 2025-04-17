package subscribe

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubscribeController) Create(ctx *gin.Context) {
	var subscribe v1.Subscribe

	if err := ctx.ShouldBindJSON(&subscribe); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	err := c.Service.Subscribes().Create(ctx, &subscribe, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, subscribe)
}
