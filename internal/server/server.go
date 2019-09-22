package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

type ServerMsg struct {
	Msg     string
	Address string
}

type Server struct {
	ErrC     chan error
	MsgC     chan ServerMsg
	Address  *net.TCPAddr
	Conn     *net.TCPConn
	Listener *net.TCPListener
}

func New(address string, port int) (*Server, error) {

	addr, err := net.ResolveTCPAddr("tcp", address+":"+strconv.Itoa(port))

	if err != nil {
		return nil, err
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		return nil, err
	}

	tcpConn, err := listener.AcceptTCP()

	if err != nil {
		log.Fatal(err)
	}

	errChannel := make(chan error, 1)
	msgChan := make(chan ServerMsg)

	server := Server{errChannel, msgChan, addr, tcpConn, listener}

	return &server, nil
}

func (server *Server) Listen() {

	go func() {
		for {
			fmt.Println("scan opened")
			scanner := bufio.NewScanner(server.Conn)

			for scanner.Scan() {
				text := scanner.Text()
				server.MsgC <- ServerMsg{text, server.Conn.RemoteAddr().String()}
			}
			fmt.Println("scan closed")

			tcpConn, err := server.Listener.AcceptTCP()

			if err != nil {
				server.ErrC <- err
				return
			}

			server.Conn = tcpConn
		}
	}()
}
