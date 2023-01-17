// Package hw11_telnet_client -- OTUS HW11 Telnet Client.
package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"
)

// TelnetClient implementation.
type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

// NewTelnetClient constructor.
func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &clientAbstraction{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type clientAbstraction struct {
	address  string
	timeout  time.Duration
	in       io.ReadCloser
	out      io.Writer
	conn     net.Conn
	inScan   *bufio.Scanner
	connScan *bufio.Scanner
}

func (t *clientAbstraction) Connect() (err error) {
	t.conn, err = net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}

	t.connScan = bufio.NewScanner(t.conn)
	t.inScan = bufio.NewScanner(t.in)

	return
}

func (t *clientAbstraction) Close() (err error) {
	if t.conn == nil {
		return
	}

	return t.conn.Close()
}

func (t *clientAbstraction) Send() (err error) {
	if t.conn == nil {
		return
	}
	if !t.inScan.Scan() {
		return errors.New("end")
	}
	_, err = t.conn.Write(append(t.inScan.Bytes(), '\n'))
	return
}

func (t *clientAbstraction) Receive() (err error) {
	if t.conn == nil {
		return
	}
	if !t.connScan.Scan() {
		return errors.New("connection closed")
	}
	_, err = t.out.Write(append(t.connScan.Bytes(), '\n'))
	return
}
