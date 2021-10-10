package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //Default返回一个默认的路由引擎
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		fmt.Println(username)
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	r.Run()
}
