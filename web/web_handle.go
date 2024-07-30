package web

import (
	"log"
	"net/http"
	"sync"

	"github.com/VladislavSCV/Test3/internal/kafka"
	"github.com/gin-gonic/gin"
)

var (
	wg      sync.WaitGroup
	topic   = "123"
	groupID = "my-group"
)

func RunWebServer() {
	r := gin.Default()
	r.LoadHTMLGlob(`C:\Users\VladislavSCV\OneDrive\Desktop\Projects\TESTS\Test3\web\templates\*`)
	r.GET("/i", GetIndexPage)
	r.GET("/ir", GetResponsePage)

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

func GetResponsePage(c *gin.Context) {
	wg.Add(1)
	var lst string
	consumer := kafka.NewKafkaConsumer([]string{"localhost:9092"}, topic, groupID)
	defer func(consumer *kafka.KafkaConsumer) {
		err := consumer.Close()
		if err != nil {

		}
	}(consumer)
	go func() {
		defer wg.Done()
		res, err := consumer.ConsumeMessages()
		if err != nil {
			log.Println("ERROR TO READ KAFKA MESSAGES")
		}
		lst = res
	}()
	wg.Wait()
	c.HTML(http.StatusOK, "respMessage.html", gin.H{
		"lst": lst,
	})
}
