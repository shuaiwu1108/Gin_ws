package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world! 我是中文",
		})
	})

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world! 我是中文Post",
		})
	})

	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.Static("/static", "./static")
	r.Run(":9000")
}
