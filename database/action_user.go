package database

import (
	"errors"
	"magic-mirror/models"
)

func IsUserExistByLogin(login string) bool {
	var count = 0 
	var user models.User
	GetDb().First(&user,models.User{Login:login}).Count(&count)
	return count ==1 && user.Id > 0
}

func Add(bean interface{}) error  {
	if !GetDb().NewRecord(bean){
		return errors.New("unable to create entity to DB")
	}
	return GetDb().Create(bean).Error
}

func GetUser(username,password string){
	var user models.User
	
}