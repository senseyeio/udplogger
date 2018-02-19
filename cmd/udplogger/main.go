package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	echoudp "github.com/senseyeio/udplogger"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	addr, err := echoudp.Addr()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("listening: ", err)
	}
	log.Printf("listening on %v\n", addr)

	conn.SetReadBuffer(1048576) // use a 1MB buffer

	go func() {
		s := <-c
		log.Printf("signal (%v) caught, terminating...", s)
		os.Exit(0)
	}()

	var buf [1024]byte
	for {
		rlen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatal("reading: ", err)
		}
		fmt.Println(string(buf[:rlen]))
	}
	os.Exit(1)
}
