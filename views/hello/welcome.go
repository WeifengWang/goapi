package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Welcome(ctx *gin.Context) {
	ctx.String(http.StatusOK, "welcome!!!")
}

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "Hello world!!!")
}

