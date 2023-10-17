package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/gorm-postgresql/controllers"
	"github.com/thitiphongD/gorm-postgresql/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.FindPosts)
	r.GET("/posts/:id", controllers.FindPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run("localhost:8080") 
}