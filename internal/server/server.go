package server

import (
	"log"
	"net"
	"strconv"
)

func StartServer(address string, port int) *net.UDPConn {

	addr, err := net.ResolveUDPAddr("udp", address+":"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf(err.Error())
	}

	listener, err := net.ListenUDP("udp", addr)

	log.Println(addr)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return listener
}
