package client

import (
	"context"
	"net"
	"strconv"
	"time"
)

type Client struct {
	Connection net.Conn
	Ctx        context.Context
	CancelCtx  context.CancelFunc
}

func New(address string, port int, timeout int64) (*Client, error) {

	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)

	conn, err := dialer.DialContext(ctx, "tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	client := Client{Connection: conn, Ctx: ctx, CancelCtx: cancel}
	return &client, nil
}
