package solution

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SolutionController) Delete(ctx *gin.Context) {
	name := ctx.Param("name")
	err := c.service.Solutions().Delete(ctx, name, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
