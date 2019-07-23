package views

import (
	"github.com/WeifengWang/goapi/views/hello"
	"github.com/gin-gonic/gin"
)


func InitRoutes(e *gin.Engine) {
	e.GET("/", hello.WebRoot)
	initV1Api(e.Group("/v1"))
}

func initV1Api(e *gin.RouterGroup) {
	e.GET("/welcome", hello.Welcome)
	e.GET("/hi", hello.SayHi)
}