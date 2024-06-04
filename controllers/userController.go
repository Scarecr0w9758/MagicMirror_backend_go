package controllers

import (
	database "magic-mirror/database"
	"magic-mirror/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// import(
// 	"context"
// 	"fmt"
// 	"log"
// 	"strconv"
// 	"net/http"
// 	"time"
// 	"time"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// helper "magic-mirror/helpers"
//  "magic-mirror/models"
//  "magic-mirror/helpers"
//  "golang.org/x/crypto/bcrypt"
// )

var validate=validator.New()

func HashPassword()

func VerifyPassword()

func GetUser() gin.HandlerFunc  {
	return func (c *gin.Context){
		userId:=c.Param("userId")
		if err:=helper.MatchUserTypeToUid(c,userId); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return

	}
	// var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)
	var user models.User
	err:=database.GetDb().First(&user,models.User{Id:userId})
	
	// defer cancel()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"errorMessage":err})
	}
	c.JSON(http.StatusOK,user)
	// .Count(&count)
	 
	}
}