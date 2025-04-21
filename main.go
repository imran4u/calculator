package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imran4u/calculator/api/handler"
)

func main() {
	router := gin.Default()

	router.GET("/add", handler.AddHandler)
	router.GET("/sub", handler.SubHandler)
	router.GET("/mul", handler.MulHandler)
	router.GET("/div", handler.DivHandler)
	router.POST("/clone", handler.CloneHandler)
	router.Run(":8080")
}
