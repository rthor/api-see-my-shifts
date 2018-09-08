package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	SetDbConn()
	defer db.Close()

	r := gin.Default()
	r.Use(AuthMiddleware())
	r.Use(CorsMiddleware())

	// will listen on the PORT env variable
	// defaults to 8080
	r.Run()
}

// CorsMiddleware ...
// Just allow everything for now... terrible
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	}
}

// AuthMiddleware ...
// Should probably read the bearer token here and throw,
// since this is the first middleware to execute in the gin pipeline
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			panic("An authorization header is present in the request... lol.")
		}
	}
}
