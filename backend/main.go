package main

import (
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {

    var router = gin.Default()
    var config = cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
        AllowHeaders:     []string{"Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return true
        },
        MaxAge: 12 * time.Hour,
    }

    router.Use(cors.New(config))
    router.GET("/ping", ping)
    router.Run()
}

func ping(context *gin.Context) {
    context.JSON(200, gin.H{
        "message": "pong",
    })
    return
}
