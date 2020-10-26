package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/someGet", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//r.GET("user/:name/*action", func(c *gin.Context) {
	//c.FullPath() == "/user/:name/*action"
	//})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastName := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastName)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.String(http.StatusOK, "ids：%v names：%v", ids, names)
	})

	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		c.SaveUploadedFile(file, "C://")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//simple group：v1
	v1 := r.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.String(http.StatusForbidden, "Forbidde")
		})
	}

	//simple group：v2
	v2 := r.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
