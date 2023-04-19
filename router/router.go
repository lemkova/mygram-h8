package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lemkova/mygram-h8/controller"
	"github.com/lemkova/mygram-h8/middlewares"

	_ "github.com/lemkova/mygram-h8/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Mygram API Documentation
// @version 1.0
// @description This is a sample server for a photo sharing app with comments feature.
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/create/", controller.CreatePhoto)
		photoRouter.GET("/", controller.GetPhotos)
		photoRouter.GET("/:id", controller.GetPhoto)
		photoRouter.GET("/user/:userId", controller.GetPhotosByUser)

		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controller.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/create/:photoId", controller.CreateComment)
		commentRouter.GET("/photo/:photoId", controller.GetComments)
		commentRouter.GET("/:commentId", controller.GetComment)

		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controller.DeleteComment)
	}
	socialMediaRouter := r.Group("/socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/create", controller.CreateSocialMedia)
		socialMediaRouter.GET("/", controller.GetAllSocialMedias)
		socialMediaRouter.GET("/:id", controller.GetSocialMedia)
		socialMediaRouter.GET("/user/:userId", controller.GetAllSocialMediasByUserID)

		socialMediaRouter.PUT("/:id", middlewares.SocialMediaAuthorization(), controller.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.SocialMediaAuthorization(), controller.DeleteSocialMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
