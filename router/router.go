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
		userRouter.PUT("/:userId", middlewares.Authentication(), middlewares.ProfileAuthorization(), controllers.UserUpdate)
		userRouter.PUT("/profile-picture", middlewares.Authentication(), middlewares.ProfileAuthorization(), controllers.UserProfilePictureUpdate)
		userRouter.DELETE("/", middlewares.Authentication(), middlewares.ProfileAuthorization(), controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.PhotoCreate)
		photoRouter.GET("/", controllers.PhotoGetAll)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoUpdate)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.PhotoDelete)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CommentCreate)
		commentRouter.GET("/", controllers.CommentList)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.CommentUpdate)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.CommentDelete)
	}

	socialmediasRouter := r.Group("/socialmedias")
	{
		socialmediasRouter.Use(middlewares.Authentication())
		socialmediasRouter.POST("/", controllers.SocialMediaCreate)
		socialmediasRouter.GET("/", controllers.SocialMediaList)
		socialmediasRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaUpdate)
		socialmediasRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.SocialMediaDelete)
	}

	return r
}
