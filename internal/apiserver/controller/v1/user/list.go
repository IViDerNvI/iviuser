package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) List(ctx *gin.Context) {

	limit, _ := strconv.Atoi(ctx.Query("limit"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))

	mapper := map[string]string{
		"username": ctx.Query("username"),
		"email":    ctx.Query("email"),
		"status":   ctx.Query("status"),
		"nickname": ctx.Query("nickname"),
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Limit:    limit,
		Offset:   offset,
		Selector: selector,
	}

	users, err := c.Srv.Users().List(ctx, listOptions)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, users)
}
