package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imran4u/calculator/calculator"
)

func SubHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter a"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter b"})
		return
	}

	calculator := &calculator.Calculator{}
	result := calculator.Subtract(a, b)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
