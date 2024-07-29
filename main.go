package main

import (
	"log"
	"sync"

	"github.com/VladislavSCV/Test3/api/rest"
	"github.com/VladislavSCV/Test3/internal/db"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go rest.RunRestServer()
	go db.RunDbServer()
	//go web.RunWebServer()
	log.Println("Hello world!")
	wg.Wait()
}
