package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	loadconfig()
	applysysctl()
	startrawdetector()
	startiptables()
	startflooddetector()
	http.HandleFunc("/", wafhandler)
	server := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(config.listenport),
	}
	go func() {
		log.Printf("axomwaf listening on port %d", config.listenport)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	<-sigchan
	server.Close()
}
