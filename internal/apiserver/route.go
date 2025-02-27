package apiserver

import (
	"github.com/gin-gonic/gin"
	ctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/user"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

func RegisterRoutes(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		user := v1.Group("/user")
		{
			userController := ctl.NewUserController(store.Factory())

			user.GET("/:id", userController.Get)
			user.POST("/", userController.Create)
			user.PUT("/:id", userController.Update)
			user.DELETE("/:id", userController.Delete)
		}
	}
}
