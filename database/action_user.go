package database

import (
	"errors"
	"magic-mirror/models"
	"net/mail"
	"strings"
)


func isStringEmail(str string) bool {
	index:=strings.Index(str,"@")
	return index > -1
}

func validateEmail(email string) bool{
	_,err:=mail.ParseAddress(email)
	return err==nil

}

func IsUserExistByLogin(login string) bool {
	var count = 0 
	var user models.User
	GetDb().First(&user,models.User{Login:login}).Count(&count)
	return count ==1 && user.Id > 0
}


func IsUserExistByEmail(email string) bool {
	var count = 0 
	var user models.User
	GetDb().First(&user,models.User{Email:email}).Count(&count)
	return count ==1 && user.Id > 0
}

func Add(bean interface{}) error  {
	if !GetDb().NewRecord(bean){
		return errors.New("unable to create entity to DB")
	}
	return GetDb().Create(bean).Error
}

func GetUser(username,password string){
	// var user models.User
	
}