package main

import (
	"fmt"
	"calculator/handler/calculator.go"
	//"calculator/handler"
	"github.com/imran4u/calculator/calculator"
)

func main() {
	c := calculator.Calculator{}
	fmt.Println(c.Add(2, 5))
	r.POST("/add", handler.AddHandler)
}
