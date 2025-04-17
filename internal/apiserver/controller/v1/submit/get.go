package submit

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (p *SubmitController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	submit, err := p.Service.Submits().Get(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, submit)
}
