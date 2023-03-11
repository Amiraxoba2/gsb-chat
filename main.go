package main

import (
	"github.com/Amiraxoba2/gsb-chat/internal/database"
	"github.com/Amiraxoba2/gsb-chat/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	database.DB.AutoMigrate(database.Chat{})
	database.DB.AutoMigrate(database.User{})
	database.DB.AutoMigrate(database.Message{})

	engine := gin.Default()
	engine.LoadHTMLGlob("resource/template/*.go.tmpl")

	engine.GET("/", route.Index)
	engine.GET("/login", route.LoginGet)
	engine.POST("/login", route.LoginPost)
	engine.GET("/output.css", func(context *gin.Context) {
		context.File("resource/template/output.css")
	})
	engine.GET("/icon", func(context *gin.Context) {
		context.File("resource/icon.svg")
	})
	engine.Run(":8080")
}
