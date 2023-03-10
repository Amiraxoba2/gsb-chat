package route

import "github.com/gin-gonic/gin"

func Login(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "ok",
		"page":   "login",
	})
}
