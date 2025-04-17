package submit

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *SubmitController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.Submits().Delete(ctx, id, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	ctx.JSON(200, nil)
}
