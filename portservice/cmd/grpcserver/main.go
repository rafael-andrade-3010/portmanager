package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"portservice/internal/core/service/portservice"
	"portservice/internal/handlers/porthandler"
	"portservice/internal/repositories/portrepo/mongo"
	"runtime"
	"sync"
	"syscall"
	"time"
)

const (
	port = ":50051"
)

func bToKb(b uint64) string {
	return fmt.Sprintf("%vKb", b/1024)
}

func PrintMemUsage() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				log.Printf("Alloc = %v, HeapAlloc = %v, TotalAlloc = %v, Sys = %v, NumGC = %v", bToKb(m.Alloc), bToKb(m.HeapAlloc), bToKb(m.TotalAlloc), bToKb(m.Sys), m.NumGC)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func main() {
	PrintMemUsage()

	repo := mongo.NewPortMongoRepo()
	service := portservice.New(repo)
	handler := porthandler.NewGrpcHandler(service)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		grpcServer.GracefulStop()
		wg.Done()
	}()

	porthandler.RegisterPortDomainServer(grpcServer, handler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
