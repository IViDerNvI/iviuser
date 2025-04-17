package solution

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SolutionController) List(ctx *gin.Context) {
	solutions, err := c.service.Solutions().List(ctx, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, solutions)
}
