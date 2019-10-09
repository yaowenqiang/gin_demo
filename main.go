package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//gin.New
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func (c *gin.Context) {
		//c.String(http.StatusOK, "Hello from %v", "Gin")
		c.HTML(http.StatusOK, "index.html", nil)
	})

	admin := r.Group("/admin")
	admin.GET("/", func (c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})
	r.Run(":3000")
	/*
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.POST("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.GET("/api/books/:isbn", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.PUT("/api/books/:isbn", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	engine.DELETE("/api/books/:isbn", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework"})
	})

	//engine.Run(port())
	*/
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	return ":" + port

}
