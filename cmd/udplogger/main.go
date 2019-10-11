package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/senseyeio/udplogger"
)

var raw bool

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logger))
	log.Println("udplogger started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-c
		log.Printf("signal (%v) caught, terminating...", s)
		os.Exit(0)
	}()

	raw = udplogger.Raw()
	addr, err := udplogger.Addr()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("listening: ", err)
	}
	log.Printf("listening on %v\n", addr)
	conn.SetReadBuffer(100 * maxUDPPayload)

	var buf [maxUDPPayload]byte
	for {
		rlen, _, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatal("reading from UDP socket: ", err)
		}
		str := string(buf[:rlen])
		if raw {
			log.Print(str)
		} else {
			log.Printf("received: %q\n", str)
		}
	}
	os.Exit(1)
}

type logger struct{}

func (_ logger) Write(b []byte) (int, error) {
	if raw {
		return fmt.Print(string(b))
	} else {
		return fmt.Print(time.Now().UTC().Format(format) + " " + string(b))
	}
}

// RFC3339 with milliseconds and right 0 padding.
const format = "2006-01-02T15:04:05.000Z07:00"

// maximum UDP payload size in bytes: 2^16 -1 −8 (UDP header) −20 (IP header)
const maxUDPPayload = 65507
