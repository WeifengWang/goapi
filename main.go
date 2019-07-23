package main

import (
	"github.com/WeifengWang/goapi/middleware"
	"github.com/WeifengWang/goapi/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//middleware
	r.Use(middleware.PrintReqAndRes)
	//init view


	views.InitRoutes(r)



	_ = r.Run(":5588")
}


