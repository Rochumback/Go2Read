package main

import (
    "database/sql"
    "fmt"
    "io"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "github.com/rochumback/Go2Read/structs"
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
    var api_group = router.Group("/api")

    api_group.POST("/create_user", create_user)
    api_group.GET("/images", get_urls)
    router.GET("/ping", ping)

    router.Run()
}

type NewUser struct {
    Username string
}

func createDatabaseConnection() *sql.DB {

    var (
        host     = os.Getenv("DB_HOST")
        port     = os.Getenv("DB_PORT")
        user     = os.Getenv("POSTGRES_USER")
        password = os.Getenv("DB_PASSWORD")
        dbname   = os.Getenv("DB_NAME")
    )
    var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var database_client, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    if err = godotenv.Load(); err != nil {
        panic(err)
    }
    return database_client
}

func sanitize() gin.HandlerFunc {
    return func(context *gin.Context) {
        body, err := io.ReadAll(context.Request.Body)
        if err != nil {
        }

        fmt.Println(len(body))
    }
}

func get_urls(context *gin.Context) {
    var images = []string{"images/3.webp"}

    context.JSON(200, gin.H{
        "images": images,
    })
}

func create_user(context *gin.Context) {
    database_client := createDatabaseConnection()

    var user = NewUser{}
    if err := context.BindJSON(&user); err != nil || user.Username == "" {
        context.JSON(400, gin.H{
            "status": "body parse error",
        })
        return
    }
    var result, err = database_client.Query(
        "SELECT * FROM manga;",
    )

    if err != nil {
        context.JSON(400, gin.H{
            "status":   err,
            "database": database_client,
        })
        return
    }

    var mangas []structs.Manga

    for result.Next() {
        manga := structs.Manga{}
        result.Scan(&manga.Id, &manga.Title, &manga.Alternative_titles, &manga.Searchable, &manga.Synopsis, &manga.Created_at, &manga.Updated_at, &manga.Created_by)
        _ = append(mangas, manga)
        fmt.Println(manga)
    }

    context.JSON(200, gin.H{
        "status": "success",
    })
    fmt.Println(mangas[0])
}

func ping(context *gin.Context) {
    context.JSON(200, gin.H{
        "message": "pong",
    })
    return
}
