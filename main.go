package main

import (
	// "fmt"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"main/config"
	"main/controller"
	"main/database"
	"main/lib/middleware"
)

func main() {
	configure := config.ConfigMain()
	db := database.InitDB()
	defer db.Close()

	// Debugging - environment variables
	// /*
	fmt.Println(configure.Server.ServerPort)
	fmt.Println(configure.Database.DbDriver)
	fmt.Println(configure.Database.DbUser)
	fmt.Println(configure.Database.DbPass)
	fmt.Println(configure.Database.DbName)
	fmt.Println(configure.Database.DbHost)
	fmt.Println(configure.Database.DbPort)
	// */

	router := SetupRouter()
	router.Run(":" + configure.Server.ServerPort)
}

// SetupRouter ...
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode

	// Write log file
	// Console color is not required to write the logs to the file
	//	gin.DisableConsoleColor()

	// Create a log file with start time
	dt := time.Now()
	t := dt.Format(time.RFC3339)
	file, _ := os.Create("start:" + t + ".log")
	gin.DefaultWriter = io.MultiWriter(file)

	// If it is required to write the logs to the file and the console
	// at the same time
	//	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// Creates a router without any middleware by default
	// router := gin.New()

	// Logger middleware: gin.DefaultWriter = os.Stdout
	// router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500
	// if there is one
	// router.Use(gin.Recovery())

	// gin.Default() = gin.New() + gin.Logger() + gin.Recovery()
	router := gin.Default()

	router.Use(middleware.CORS())

	// API:v1.0
	// Non-protected routes
	v1 := router.Group("/api/v1/")
	{
		// User
		v1.GET("users", controller.GetUsers)
		v1.GET("users/:id", controller.GetUser)
		v1.POST("users", controller.CreateUser)
		v1.PUT("users/:id", controller.UpdateUser)
		v1.DELETE("users/:id", controller.DeleteUser)

		// Post
		v1.GET("posts", controller.GetPosts)
		v1.GET("posts/:id", controller.GetPost)
		v1.POST("posts", controller.CreatePost)
		v1.PUT("posts/:id", controller.UpdatePost)
		v1.DELETE("posts/:id", controller.DeletePost)
	}

	return router
}
