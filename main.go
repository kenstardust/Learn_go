package main

import (
	"Kee"
	"net/http"
)

func main() {

	r := Kee.New()

	r.GET("/", func(c *Kee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *Kee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *Kee.Context) {
			c.String(http.StatusOK, "Hello %s! your path is %s", c.Param("name"), c.Path)
		})
	}

	r.POST("/login", func(c *Kee.Context) {
		c.JSON(http.StatusOK, Kee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")

}
