package route

import (
	"github.com/Amiraxoba2/gsb-chat/internal/database"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	var chats []database.Chat
	database.DB.Find(&chats)
	context.HTML(200, "index.go.tmpl", gin.H{
		"Chats": chats,
	})
}
