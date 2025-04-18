package like

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *LikeController) List(ctx *gin.Context) {
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	mapper := map[string]string{
		"item_type": ctx.Param("type"),
		"username":  ctx.Param("id"),
	}

	selector := v1.Selector(mapper)

	result, err := c.Service.Likes().List(ctx, &v1.ListOptions{
		Offset:   offset,
		Limit:    limit,
		Selector: selector,
	})
	if err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseQuery, nil)
		return
	}

	core.WriteResponse(ctx, nil, result)
}
