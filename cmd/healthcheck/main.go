// Program healthcheck returns 0 if there is already a process listening
// on localhost at the UDP port specified by the environment variable
// PORT. It returns 1 if the port is free or 2 if there is another error.
package main

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/senseyeio/udplogger"
)

const (
	portEnvVar = "PORT"
)

func main() {
	addr, err := udplogger.Addr()
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err == nil {
		_ = conn.Close()
		os.Exit(1)
	}
	if !strings.HasSuffix(err.Error(), "address already in use") {
		log.Println(err)
		os.Exit(2)
	}
	os.Exit(0)
}
