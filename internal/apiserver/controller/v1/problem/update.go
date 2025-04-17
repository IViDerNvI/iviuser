package problem

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *ProblemController) Update(ctx *gin.Context) {
	var p v1.Problem
	if err := ctx.ShouldBindJSON(&p); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	err := c.Service.Problems().Update(ctx, &p, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, p)
}
