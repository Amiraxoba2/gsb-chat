package main

import (
	"github.com/Amiraxoba2/gsb-chat/internal/database"
	"github.com/Amiraxoba2/gsb-chat/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.DB.AutoMigrate(database.Chat{})
	if err != nil {
		return
	}
	engine := gin.Default()
	engine.LoadHTMLGlob("resource/template/*.go.tmpl")
	engine.GET("/", route.Index)
	engine.GET("/login", route.Login)
	engine.GET("/output.css", func(context *gin.Context) {
		context.File("resource/template/output.css")
	})
	engine.GET("/icon", func(context *gin.Context) {
		context.File("resource/icon.svg")
	})
	engine.Run(":8080")
}
