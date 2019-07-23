package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SayHi(ctx *gin.Context) {
	name := ctx.Query("name")

	if name != "" {
		ctx.String(http.StatusOK, "hi %s!", name)
	} else {
		ctx.String(http.StatusOK, "Hello please tell me your name!")
	}
}
