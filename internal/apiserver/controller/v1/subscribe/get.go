package subscribe

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubscribeController) Get(ctx *gin.Context) {
	rsrcId := ctx.Param("resourceid")
	rsrcType := ctx.Param("type")

	sub := &v1.Subscribe{
		ItemType: rsrcType,
		ItemID:   rsrcId,
	}

	result, err := c.Service.Subscribes().Get(ctx, sub, &v1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, result)
}
