package database

import (
	"fmt"
	"log"
	"time"

	"magic-mirror/models"

	"github.com/jinzhu/gorm"
)

var dbase *gorm.DB
// Функция с большой буквы - публичная, с маленькой - приватная!
func Init() *gorm.DB{
	db, err:=gorm.Open("postgres","user=admin password=admin dbname=postgres sslmode=disable")
	
	if err!=nil{
		log.Fatal("Error while open coonection to DB",err)
	}

	db.AutoMigrate(&models.Tokens{},&models.User{},&models.Task{})
	return db

}

func GetDb() *gorm.DB {
	if dbase==nil{
		dbase=Init()
		var sleep=time.Duration(1)
		for dbase==nil {
			sleep=sleep*2
			fmt.Printf("DB is not ready, sleeping for %d seconds\n",sleep)
			time.Sleep(sleep*time.Second)
			dbase=Init()
		}
	}
	return dbase
}