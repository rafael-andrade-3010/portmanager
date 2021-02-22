package main

import (
	"fmt"
	"log"
	"net/http"
	"portapi/api"
	"runtime"
	"time"
)

func setupRoutes() {
	api.SetupRoutes()
}

func bToKb(b uint64) string {
	return fmt.Sprintf("%vKb", b / 1024)
}

func PrintMemUsage() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				log.Printf("Alloc = %v, HeapAlloc = %v, TotalAlloc = %v, Sys = %v, NumGC = %v", bToKb(m.Alloc), bToKb(m.HeapAlloc	), bToKb(m.TotalAlloc), bToKb(m.Sys), m.NumGC)
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func main() {
	PrintMemUsage()
	fmt.Println("Starting server")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
