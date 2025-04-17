package submit

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubmitController) Create(ctx *gin.Context) {
	var submit v1.Submit

	if err := ctx.ShouldBindJSON(&submit); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Submits().Create(ctx, &submit, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	ctx.JSON(200, submit)
}
