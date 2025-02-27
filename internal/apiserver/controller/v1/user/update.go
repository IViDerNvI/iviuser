package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/iviuser/model/v1"
	"github.com/ividernvi/iviuser/pkg/core"
)

func (c *UserController) Update(gin *gin.Context) {
	var user v1.User
	if err := gin.ShouldBindJSON(&user); err != nil {
		core.WriteResponse(gin, err, nil)
		return
	}
	if err := c.Srv.Users().Update(gin, &user, nil); err != nil {
		core.WriteResponse(gin, err, nil)
		return
	}
	gin.JSON(http.StatusOK, user)
}
