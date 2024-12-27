package main

import "github.com/gin-gonic/gin"

func main() {
	// c := calculator.Calculator{}
	// fmt.Println(c.Add(2, 5))
	r := gin.Default()

	r.GET("/add" AddHandler)

}
