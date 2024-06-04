package main

import (
	"fmt"
	"magic-mirror/database"
	"magic-mirror/handlers"
	routes "magic-mirror/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
 
func main() {
    port:=os.Getenv("PORT")

    if port==""{
        port="8000"
    }   

    database.Init()
    gin.SetMode(gin.ReleaseMode)
    router:=gin.Default()
   
    routes.WidgetRoutes(router)
    routes.UserRoutes(router)
    router.Use(cors.Default())
    router.GET("/hello",handlers.HelloWorld)
    router.POST("/user/registration",handlers.Registration)
    // router.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"https://foo.com"},
    //     AllowMethods:     []string{"PUT", "PATCH","POST"},
    //     AllowHeaders:     []string{"Origin"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     AllowOriginFunc: func(origin string) bool {
    //         return origin == "https://github.com"
    //     },
    //     MaxAge: 12 * time.Hour,
    // }))

    router.Run(":8080")
    fmt.Println("Server is running! on Port")
}