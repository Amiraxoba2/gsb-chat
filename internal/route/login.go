package route

import (
	"github.com/Amiraxoba2/gsb-chat/internal/database"
	"github.com/gin-gonic/gin"
)

func LoginGet(context *gin.Context) {
	context.HTML(200, "login.go.tmpl", gin.H{})
}

func LoginPost(context *gin.Context) {
	name := context.PostForm("username")
	password := context.PostForm("password")
	var user database.User
	database.DB.First(&user, "name = ?", name)
	if user.Password == password {
		context.Writer.WriteString("")
	} else {
		context.Redirect(302, "/login")
	}
}
