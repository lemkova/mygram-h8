package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lemkova/mygram-h8/controller"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
	}
	return r
}
