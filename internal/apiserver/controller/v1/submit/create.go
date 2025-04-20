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

	operatorName, exists := ctx.Get("X-Operation-User-Name")
	if !exists {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	var valid bool
	submit.Author, valid = operatorName.(string)
	if !valid {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	submit.Status = v1.SubmitStatusPending

	if err := submit.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Submits().Create(ctx, &submit, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, submit)
}
