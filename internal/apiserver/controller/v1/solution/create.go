package solution

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SolutionController) Create(ctx *gin.Context) {
	var solution v1.Solution
	if err := ctx.ShouldBindJSON(&solution); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	err := c.service.Solutions().Create(ctx, &solution, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	core.WriteResponse(ctx, nil, solution)
}
