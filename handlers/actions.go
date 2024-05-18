package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"response": gin.H{
			"httpMethod": http.MethodGet,
			"code": http.StatusOK,
			"message": "hello World!",
		},
	})
}