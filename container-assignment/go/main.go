package main

import "github.com/gin-gonic/gin"

// main function
func main() {
	r := gin.Default()
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, from endpoint 1",
		})
	})

	r.GET("/endpoint2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, fron endpoint 2!",
		})
	})

	r.Run()
}
