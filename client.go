package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

var ErrNotOpen = errors.New("connection not open")

type Client struct {
	Timeout time.Duration
	Address string
	conn    net.Conn
	reader  *bufio.Reader
}

func (c *Client) Open() (err error) {
	c.conn, err = net.DialTimeout("tcp", c.Address, c.Timeout)
	if err != nil {
		return
	}
	c.reader = bufio.NewReader(c.conn)

	return
}

func (c *Client) Close() error {
	if c.conn == nil {
		return ErrNotOpen
	}
	return c.conn.Close()
}

func (c *Client) Exec(cmd string) (string, error) {
	if c.conn == nil {
		return "", ErrNotOpen
	}

	_, err := fmt.Fprintf(c.conn, cmd+"\r\n")
	if err != nil {
		return "", err
	}

	return c.reader.ReadString('\n')
}
