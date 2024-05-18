package main

import (
	"fmt"
	"magic-mirror/database"
	"magic-mirror/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
 
func main() {
    database.Init()
    gin.SetMode(gin.ReleaseMode)
    router:=gin.Default()
    router.GET("/hello",handlers.HelloWorld)


    router.Run(":8080")
    fmt.Println("Server is running! on Port")
}