package main

import (
	"log"
	"sync"

	"github.com/VladislavSCV/Test3/api/rest"
	"github.com/VladislavSCV/Test3/internal/db"
	"github.com/VladislavSCV/Test3/web"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(4)
	go rest.RunRestServer()
	go db.RunDbServer()
	go web.RunWebServer()
	log.Println("Hello world!")
	wg.Wait()
}
