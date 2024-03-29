package main

import (
	"os"
)

func main() {

	r := RegisterRoutes()
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

	*/
	r.Run("3000")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	return ":" + port

}
