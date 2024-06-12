package handlers

import (
	"encoding/json"
	"magic-mirror/database"
	"magic-mirror/models"
	"net/http"
	"net/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

func isStringEmail(str string) bool {
	index:=strings.Index(str,"@")
	return index > -1
}

func isValidEmail(email string) bool{
	_,err:=mail.ParseAddress(email)
	return err==nil

}

func HelloWorld(ctx *gin.Context){
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
func Registration(context *gin.Context){
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var user *models.User

	decode :=json.NewDecoder(context.Request.Body).Decode(&user)
	
	if decode != nil{
		context.JSON(http.StatusOK,gin.H{
			"response":decode.Error(),
		})
		return
	}


	isExistByLogin := database.IsUserExistByLogin(user.Login)
	if !isExistByLogin{
		user:= &models.User{
			Id:user.Id,
			Login:user.Login,
			Email:user.Email,
			Name:user.Name,
			Password: user.Password,
			}
		isSuccessAddError:=database.Add(user)		
		if isSuccessAddError==nil{
			context.JSON(http.StatusOK,gin.H{
				"response":gin.H{
					"code":http.StatusOK,
					"message":"Вы успешно зарегистрированы, используйте логин для входа!",
				},
			})
		}
	} else{
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"errorMessage":"Пользователь с таким логином существует",
		})
	}
}



func SignIn(context *gin.Context){
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	var user *models.User

	decode :=json.NewDecoder(context.Request.Body).Decode(&user)

	if decode != nil{
		context.JSON(http.StatusOK,gin.H{
			"response":decode.Error(),
		})
		return
	}

	// Проверям нам прислали почту или логин
	if isStringEmail(user.Login){
		if isValidEmail(user.Login){
			isExistByEmail := database.IsUserExistByEmail(user.Login)
			if isExistByEmail{
				context.JSON(http.StatusOK,gin.H{
				"response":gin.H{
					"message":"Вы успешно вошли, используя почту!",
					"canEnter":true,
				},
				})
			} else{
				context.JSON(http.StatusUnprocessableEntity,gin.H{
					"canEnter":false,
					"errorMessage":"Почта не найдена",

				})
			}
		}
	} else {
	isExistByLogin := database.IsUserExistByLogin(user.Login)
		// context.JSON(http.StatusOK,gin.H{
		// 		"response":gin.H{
		// 			"code":database.GetDb().First(&user,models.User{Login:user.Login}),
		// 			"getDb":database.GetDb(),
		// 			"User":&user,
		// 		},
		// 	})
	if isExistByLogin {
		context.JSON(http.StatusOK,gin.H{
		"response":gin.H{
			"message":"Вы успешно вошли по логину!",
			"canEnter":true,
		},
		})
	
	} else{
		context.JSON(http.StatusUnprocessableEntity,gin.H{
			"canEnter":false,
			"errorMessage":"Логин не найден",
		})
	}
}}