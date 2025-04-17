package solution

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SolutionController) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	solution, err := c.service.Solutions().Get(ctx, name, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, solution)
}
