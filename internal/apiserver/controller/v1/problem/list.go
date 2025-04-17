package problem

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *ProblemController) List(ctx *gin.Context) {
	problems, err := c.Service.Problems().List(ctx, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, problems)
}
