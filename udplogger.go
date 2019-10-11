package udplogger

import (
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"strconv"
)

func Addr() (*net.UDPAddr, error) {
	s, ok := os.LookupEnv("PORT")
	if !ok {
		return nil, errors.New("missing PORT environment variable")
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, fmt.Errorf("parsing PORT environment variable: %v", err)
	}
	if i < 0 {
		return nil, fmt.Errorf("invalid port: negative (%d)", i)
	}
	if i > math.MaxUint16 {
		return nil, fmt.Errorf("invalid port: too big (%d)", i)
	}

	return &net.UDPAddr{
		Port: i,
		IP:   net.ParseIP("0.0.0.0"),
	}, nil
}

func Raw() bool {
	val, ok := os.LookupEnv("RAW")
	if !ok {
		return false
	}
	return val == "on"
}
