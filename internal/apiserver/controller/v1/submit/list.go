package submit

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubmitController) List(ctx *gin.Context) {
	submits, err := c.Service.Submits().List(ctx, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, submits)
}
