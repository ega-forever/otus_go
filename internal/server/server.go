package server

import (
	"net"
	"strconv"
	"strings"
)

type ServerMsg struct {
	Msg     string
	Address string
}

type Server struct {
	ErrC     chan error
	MsgC     chan ServerMsg
	Address  *net.UDPAddr
	Listener *net.UDPConn
}

func New(address string, port int) (*Server, error) {

	addr, err := net.ResolveUDPAddr("udp", address+":"+strconv.Itoa(port))

	if err != nil {
		return nil, err
	}

	listener, err := net.ListenUDP("udp", addr)

	if err != nil {
		return nil, err
	}

	errChannel := make(chan error, 1)
	msgChan := make(chan ServerMsg)

	server := Server{errChannel, msgChan, addr, listener}

	return &server, nil
}

func (server *Server) Listen() {

	go func() {
		msg := make([]byte, 1024)
		for {
			_, fromAddress, err := server.Listener.ReadFromUDP(msg)

			msgs := strings.Split(strings.Trim(string(msg), "\x00"), "\n")

			isPing := false
			userMessage := ""

			for _, msgStr := range msgs {

				if msgStr == "ping" {
					isPing = true
				} else if len(msgStr) > 0 {
					userMessage = msgStr
				}
			}

			if err == nil && len(userMessage) > 0 {
				server.MsgC <- ServerMsg{userMessage, fromAddress.String()}
			}

			if isPing {
				_, _ = server.Listener.WriteToUDP([]byte("pong"), fromAddress)
			}

			msg = make([]byte, 1024)
		}
	}()
}
