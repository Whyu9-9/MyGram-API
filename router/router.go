package router

import (
	"mygram-api/controllers"
	"mygram-api/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:id", middlewares.Authentication(), middlewares.ProfileAuthorization(), controllers.UserUpdate)
		userRouter.DELETE("/:id", middlewares.Authentication(), middlewares.ProfileAuthorization(), controllers.UserDelete)
	}

	return r
}
