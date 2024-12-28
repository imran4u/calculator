package handler

import (
	"net/http"
	"strconv"

	"github.com/imran4u/calculator/calculator"

	"github.com/gin-gonic/gin"
)

func DivHandler(c *gin.Context) {
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
	result, err := calculator.Div(a, b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
