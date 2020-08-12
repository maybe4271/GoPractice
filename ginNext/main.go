package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Hello before")

		c.Abort()

		fmt.Println("Hello end")
	}
}

func SessionMiddleware() gin.HandlerFunc {
	return func(c * gin.Context) {
		fmt.Println("World before")
		c.Next()

		fmt.Println("World end")
	} 
}


func main() {
	fmt.Println("Hello World")
	r := gin.Default()

	r.Use(AuthMiddleware())
	r.Use(SessionMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
