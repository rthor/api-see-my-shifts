package main

import (
	"github.com/gin-gonic/gin"
)

func MainRouteHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "Message": "Works",
    })
}
