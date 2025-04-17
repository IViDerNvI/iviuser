package comment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *CommentController) Create(ctx *gin.Context) {
	var comment v1.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	opUserName := ctx.MustGet("X-Operation-User-Name").(string)
	opUserStatus := ctx.MustGet("X-Operation-User-Status").(string)

	if opUserStatus != "admin" && opUserName != comment.Auhtor {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := comment.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Comments().Create(ctx, &comment, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	ctx.JSON(200, comment)
}
