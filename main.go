package main

import (
	"github.com/WeifengWang/goapi/views"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//middleware
	//init view


	views.InitRoutes(r)



	_ = r.Run(":5588")
}


