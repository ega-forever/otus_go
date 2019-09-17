package client

import (
	"net"
	"strconv"
	"time"
)

type Client struct {
	ErrC       chan error
	Address    *net.UDPAddr
	Connection *net.UDPConn
}

func New(address string, port int) (*Client, error) {

	serverAddr, err := net.ResolveUDPAddr("udp", address+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return nil, err
	}

	errChannel := make(chan error, 1)
	client := Client{errChannel, serverAddr, conn}

	return &client, nil
}

func (client *Client) Ping(milliseconds time.Duration) {

	go func() {

		for {

			timer := time.NewTimer(milliseconds * time.Millisecond)
			<-timer.C

			err := client.Connection.SetReadDeadline(time.Now().Add(time.Second))

			if err != nil {
				client.ErrC <- err
				return
			}

			_, err = client.Connection.Write([]byte("ping\n"))

			if err != nil {
				client.ErrC <- err
				return
			}

			msg := make([]byte, 1024)
			_, _, err = client.Connection.ReadFromUDP(msg)

			if err != nil {
				client.ErrC <- err
			}
		}
	}()

}
