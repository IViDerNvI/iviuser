package comment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *CommentController) Update(ctx *gin.Context) {
	var comment v1.Comment
	if err := ctx.BindJSON(&comment); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	com, err := c.Service.Comments().Get(ctx, comment.InstanceID, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}
	opUserName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserStatus, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatus == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	if opUserStatus != "admin" && opUserName != com.Auhtor {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := c.Service.Comments().Update(ctx, &comment, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseUpdate, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
