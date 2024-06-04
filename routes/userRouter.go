package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authorization())

	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:userId",controller.GetUser())


	// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	ctx.JSON(http.StatusOK,gin.H{
		"response": gin.H{
			"httpMethod": http.MethodGet,
			"code": http.StatusOK,
			"message": "hello World!",
		},
	})
}