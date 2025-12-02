package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "OK",
	})

}
