package core

import "github.com/gin-gonic/gin"

func WriteResponse(ctx *gin.Context, err error, data interface{}) {
	if err != nil {
		if customErr, ok := err.(interface{ ErrCode() int }); ok {
			ctx.JSON(customErr.ErrCode(), gin.H{
				"code":    customErr.ErrCode(),
				"status":  "error",
				"message": err.Error(),
				"data":    data,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
			"data":    data,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"status":  "success",
		"message": "success",
		"data":    data,
	})
}
