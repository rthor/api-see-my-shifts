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

    // Migrate the schema
    db.AutoMigrate(&User{}, &Team{}, &TeamPlan{}, &UserTeam{}, &Shift{}, &Day{})

    r.GET("api/users", func(c *gin.Context) {
        res := []User{}
        db.Find(&res)

		c.JSON(200, gin.H{
			"data": res,
		})
	})

    r.DELETE("api/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        user := User{}
        db.First(&user, id)
        db.Delete(&user)

		c.JSON(200, gin.H{
            "message": "User deleted successfully",
			"data": user,
		})
	})

    r.POST("api/shifts", func(c *gin.Context) {
		id := c.PostForm("id")
        user := User{}
        db.First(&user, id)

		c.JSON(200, gin.H{
			"data": user,
		})
	})

    r.GET("api/days", func(c *gin.Context) {
        res := []Day{}
        db.Find(&res)

		c.JSON(200, gin.H{
			"data": res,
		})
	})

    r.POST("api/users", func(c *gin.Context) {
        var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

        db.Create(&user)

        c.JSON(201, gin.H{
            "data": user,
        })
    })

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
