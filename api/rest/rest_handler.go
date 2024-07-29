package rest

import (
	"log"
	"net/http"
	"sync"

	"github.com/VladislavSCV/Test3/internal/db"

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
	r.POST("/saveMessage", SaveMessage)

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

func SaveMessage(c *gin.Context) {
	var message RequestMessage
	if c.ShouldBind(&message) != nil {
		log.Println(message.Message)
	}
	db.AddMessageToDB(message.Message)
	c.Redirect(http.StatusFound, "http://localhost:8080/i")
}
