package main

// import "github.com/gin-gonic/gin"

import (
	"log"
	AuthControler "melivecode/jwt-api/controllers/auth"
	UserControler "melivecode/jwt-api/controllers/user"
	middleware "melivecode/jwt-api/middleware"
	orm "melivecode/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading Env File")
	}
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())

	// var f orm.User
	// var auth88 auth.RegisterBody

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", AuthControler.Register)
	r.POST("/login", AuthControler.Login)

	authorized := r.Group("/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserControler.ReadAll)
	authorized.GET("/profile", UserControler.Profile)

	r.Run("localhost:8000")

}

// go mod init melivecode
// go get . เพื่อ d/l ไฟล์ lib
// go mod tidy
// go get github.com/gin-gonic/gin
