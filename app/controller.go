package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes() *gin.Engine {

	r := gin.Default()
	//gin.New
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func (c *gin.Context) {
		//c.String(http.StatusOK, "Hello from %v", "Gin")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, id, nil)
	})

	//http --auth admin:admin --auth-type basic "localhost:3001/admin/"
	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})


	//r.Static("/static", "./static")
	r.StaticFS("/static", http.Dir("static"))

	//r.StaticFS("/publc",http.Dir("./public"))
	r.Run(":3001")
	return r
}

