package handler

import (
	"net/http"
	"strconv"

	"github.com/imran4u/calculator/api/model"
	"github.com/imran4u/calculator/calculator"

	"github.com/gin-gonic/gin"
)

func AddHandler(c *gin.Context) {
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
	result := calculator.Add(a, b)
	r := model.AddResult{
		BaseResult: model.BaseResult{First: a, Second: b},
		Result:     result,
	}

	c.JSON(http.StatusOK, r)
}
