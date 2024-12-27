package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {
	//TODO: Fix the func for addition
	// PARSE THE REQUEST A, B,
	// dO THE ADD FUCTION
	// RETURN RESULT

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})

	// {
	// 	A: 2,
	// 	B: 3,
	// 	ADD:5
	// }
}
