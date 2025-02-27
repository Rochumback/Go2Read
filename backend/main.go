package main

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/gin-gonic/gin"
)

func main() {

    var (
        host     = os.Getenv("DB_HOST")
        port     = os.Getenv("DB_PORT")
        user     = os.Getenv("POSTGRES_USER")
        password = os.Getenv("DB_PASSWORD")
        dbname   = os.Getenv("DB_NAME")
    )

    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    database, error := sql.Open("postgres", psqlInfo)
    if error != nil {
        panic(error)
    }
    connection, error := database.Driver().Open("uwu")
    if error != nil {
        panic(error)
    }

    router := gin.Default()
    router.GET("/ping", ping)

    router.Run()
}

func ping(context *gin.Context) {
    context.JSON(200, gin.H{
        "message": "pong",
    })
    return
}
