package apiserver

import (
	"github.com/gin-gonic/gin"
	commctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/comment"
	likectl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/like"
	postctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/post"
	probctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/problem"
	soluctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/solution"
	sbmtctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/submit"
	sbscctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/subscribe"
	userctl "github.com/ividernvi/iviuser/internal/apiserver/controller/v1/user"
	usermiddleware "github.com/ividernvi/iviuser/internal/apiserver/middlewares/user"
	"github.com/ividernvi/iviuser/internal/apiserver/store"
)

func RegisterRoutes(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		// controllers
		userController := userctl.NewUserController(store.Factory())
		likeController := likectl.NewLikeController(store.Factory())
		postController := postctl.NewPostController(store.Factory())
		commentController := commctl.NewCommentController(store.Factory())
		problemController := probctl.NewProblemController(store.Factory())
		submitController := sbmtctl.NewSubmitController(store.Factory())
		subscribeController := sbscctl.NewSubscribeController(store.Factory())
		solutionController := soluctl.NewSolutionController(store.Factory())

		// middlewares
		authorize := usermiddleware.Authorize(userController)
		mustLogin := usermiddleware.MustLogin()
		mustAdmin := usermiddleware.MustAdmin()

		v1.POST("/login", authorize, mustLogin, userController.Login)
		v1.POST("/logout", userController.Logout)
		v1.POST("/refresh", authorize, userController.Refresh)

		user := v1.Group("/user")
		{
			user.GET("/:id", authorize, userController.Get)
			user.POST("/", authorize, userController.Create)
			user.PUT("/:id", authorize, mustLogin, userController.Update)
			user.DELETE("/:id", authorize, mustLogin, userController.Delete)
			user.GET("/", authorize, userController.List)

			// user like items
			user.GET("/:id/:type/like", likeController.List)

			// user subscribe items
			user.GET("/:id/:type/subscribe", subscribeController.List)

			// user avator
			user.GET("/:id/avatar", userController.GetAvatar)
			user.PUT("/:id/avatar", authorize, mustLogin, userController.PutAvatar)
		}

		post := v1.Group("/post")
		{
			post.GET("/:id", authorize, postController.Get)
			post.POST("/", authorize, mustLogin, postController.Create)
			post.PUT("/:id", authorize, mustLogin, postController.Update)
			post.DELETE("/:id", authorize, mustLogin, postController.Delete)
			post.GET("/", authorize, postController.List)

			// post comment
			post.GET("/:id/comment/:commentid", authorize, commentController.Get)
			post.GET("/:id/comment/", authorize, commentController.List)
			post.POST("/:id/comment/", authorize, mustLogin, commentController.Create)
			post.PUT("/:id/comment/:commentid", authorize, mustLogin, commentController.Update)
			post.DELETE("/:id/comment/:commentid", authorize, mustLogin, commentController.Delete)
		}

		like := v1.Group("/like")
		{
			like.POST("/:type/:resourceid", authorize, mustLogin, likeController.Create)
			like.GET("/:type/:resourceid", authorize, likeController.Get)
			like.DELETE("/:type/:resourceid", authorize, mustLogin, likeController.Delete)
		}

		subscribe := v1.Group("/subscribe")
		{
			subscribe.POST("/:type/:resourceid", authorize, mustLogin, subscribeController.Create)
			subscribe.GET("/:type/:resourceid", authorize, subscribeController.Get)
			subscribe.DELETE("/:type/:resourceid", authorize, mustLogin, subscribeController.Delete)
		}

		problem := v1.Group("/problem")
		{
			problem.GET("/:id", authorize, problemController.Get)
			problem.POST("/", authorize, mustAdmin, problemController.Create)
			problem.PUT("/:id", authorize, mustAdmin, problemController.Update)
			problem.DELETE("/:id", authorize, mustAdmin, problemController.Delete)
			problem.GET("/", authorize, problemController.List)
		}

		submit := v1.Group("/submit")
		{
			submit.GET("/:id", authorize, submitController.Get)
			submit.POST("/", authorize, mustLogin, submitController.Create)
			submit.PUT("/:id", authorize, mustAdmin, submitController.Update)
			submit.DELETE("/:id", authorize, mustAdmin, submitController.Delete)
			submit.GET("/", authorize, submitController.List)
		}

		solution := v1.Group("/solution")
		{
			solution.GET("/:id", authorize, solutionController.Get)
			solution.POST("/", authorize, solutionController.Create)
			solution.PUT("/:id", authorize, solutionController.Update)
			solution.DELETE("/:id", authorize, solutionController.Delete)
			solution.GET("/", authorize, solutionController.List)
		}

	}
}
