package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunWebServer() {
	r := gin.Default()
	r.LoadHTMLGlob(`C:\Users\VladislavSCV\OneDrive\Desktop\Projects\TESTS\Test3\web\templates\*`)
	r.GET("/i", GetIndexPage)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func GetIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
