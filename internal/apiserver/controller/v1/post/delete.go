package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *PostController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	post, err := c.Service.Posts().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName := ctx.MustGet("X-Operation-User-Name").(string)
	operatorStatus := ctx.MustGet("X-Operation-User-Status").(string)

	if operatorStatus != "admin" && operatorName != post.Author {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := c.Service.Posts().Delete(ctx, uint(id), nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
