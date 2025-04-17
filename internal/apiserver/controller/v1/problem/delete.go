package problem

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *ProblemController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.Problems().Delete(ctx, id, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
