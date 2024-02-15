package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// 启动Gin应用程序
	r.Run()
}
