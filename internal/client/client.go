package client

import (
	"log"
	"net"
	"strconv"
)

func StartClient(address string, port int) *net.UDPConn {

	serverAddr, err := net.ResolveUDPAddr("udp", address+":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("cannot resolve server addr: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("cannot dial to server: %v", err)
	}

	return conn
}
