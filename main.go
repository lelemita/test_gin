package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lelemita/test_gin/database"
	"github.com/lelemita/test_gin/entities"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	user := entities.NewUser(1, "Anne")
	result := database.DB.Create(&user)
	fmt.Println(result)

	r := setupRouter()
	r.Run(":8080")
}
