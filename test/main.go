package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "傻逼",
	})
}
func main() {
	r := gin.New()
	r.GET("/test", test)
	r.Run(":8023")
}
