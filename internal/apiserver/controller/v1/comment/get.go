package comment

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *CommentController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	comment, err := c.Service.Comments().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, comment)
}
