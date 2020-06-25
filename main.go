package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/markshao/redme/server"
	"golang.org/x/net/context"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var chanConn = make(chan net.Conn, 10)
	go func() {
		server.StartRedmeServer(ctx, chanConn)
	}()

	log.Println("handle request")

	for conn := range chanConn {
		pool.Process(conn)
	}

	// singal handler
	signals := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		signal := <-signals
		log.Println(signal)
		cancel()
		done <- struct{}{}
	}()

	log.Println("Server Started")
	<-done

}
