package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *PostController) List(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	mapper := map[string]string{
		"title":   ctx.Query("title"),
		"content": ctx.Query("content"),
		"author":  ctx.Query("author"),
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Limit:    limit,
		Offset:   offset,
		Selector: selector,
	}

	posts, err := c.Service.Posts().List(ctx, listOptions)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, posts)
}
