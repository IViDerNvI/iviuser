package comment

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *CommentController) Delete(ctx *gin.Context) {
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

	opUserName := ctx.MustGet("X-Operation-User-Name").(string)
	opUserStatus := ctx.MustGet("X-Operation-User-Status").(string)

	if opUserStatus != "admin" && opUserName != comment.Auhtor {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := c.Service.Posts().Delete(ctx, uint(id), nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
