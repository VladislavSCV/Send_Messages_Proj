package rest

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	wg sync.WaitGroup
)

type RequestMessage struct {
	Message string `form:"message"`
}

func RunRestServer() {
	defer wg.Done()
	r := gin.Default()

	r.GET("/ping", Pong)
	r.POST("/sendMessage", SendMessage)

	err := r.Run(":8000")
	if err != nil {
		return
	}
}

func Pong(c *gin.Context) {
	log.Println("GOOD")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func SendMessage(c *gin.Context) {
	var message RequestMessage
	if c.ShouldBind(&message) != nil {
		log.Println(message.Message)
	}
	c.String(http.StatusOK, "OK")
}
