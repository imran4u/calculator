package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imran4u/calculator/api/model"
	"github.com/imran4u/calculator/calculator"
)

func SubHandler(c *gin.Context) {
	aStr := c.Query("a")
	bStr := c.Query("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid a value, it should be integer"})
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid b value, it should be integer"})
		return
	}

	calculator := &calculator.Calculator{}
	result := calculator.Subtract(a, b)
	r := model.SubResult{
		BaseResult: model.BaseResult{First: a, Second: b},
		Result:     result,
	}
	c.JSON(http.StatusOK, r)
}
