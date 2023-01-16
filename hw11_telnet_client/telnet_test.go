package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTelnetClient(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")
		require.NoError(t, err)
		defer func() { require.NoError(t, l.Close()) }()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			in := &bytes.Buffer{}
			out := &bytes.Buffer{}

			timeout, err := time.ParseDuration("10s")
			require.NoError(t, err)

			client := NewTelnetClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
			require.NoError(t, client.Connect())
			defer func() { require.NoError(t, client.Close()) }()

			in.WriteString("hello\n")
			err = client.Send()
			require.NoError(t, err)

			err = client.Receive()
			require.NoError(t, err)
			require.Equal(t, "world\n", out.String())
		}()

		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			require.NoError(t, err)
			require.NotNil(t, conn)
			defer func() { require.NoError(t, conn.Close()) }()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			require.NoError(t, err)
			require.Equal(t, "hello\n", string(request)[:n])

			n, err = conn.Write([]byte("world\n"))
			require.NoError(t, err)
			require.NotEqual(t, 0, n)
		}()

		wg.Wait()
	})
}

func testRun(t *testing.T, wg *sync.WaitGroup, l net.Listener) {
	t.Helper()
	go func() {
		defer wg.Done()

		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout, err := time.ParseDuration("60s")
		require.NoError(t, err)

		client := NewTelnetClient(l.Addr().String(), timeout, io.NopCloser(in), out)
		require.NoError(t, client.Connect())
		defer func() { require.NoError(t, client.Close()) }()

		in.WriteString("HELLO\n")
		err = client.Send()
		require.NoError(t, err)

		err = client.Receive()
		require.NoError(t, err)
		require.Equal(t, "WORLD!\n", out.String())
	}()
}

func testCheckConnect(t *testing.T, wg *sync.WaitGroup, l net.Listener) {
	t.Helper()
	defer wg.Done()

	conn, err := l.Accept()
	require.NoError(t, err)
	require.NotNil(t, conn)
	defer func() { require.NoError(t, conn.Close()) }()

	request := make([]byte, 1024)
	n, err := conn.Read(request)
	require.NoError(t, err)
	require.Equal(t, "HELLO\n", string(request)[:n])

	n, err = conn.Write([]byte("WORLD!\n"))
	require.NoError(t, err)
	require.NotEqual(t, 0, n)
}

func testIsCloseServer(t *testing.T, wg *sync.WaitGroup, clientWg *sync.WaitGroup, l net.Listener) {
	t.Helper()
	defer func() {
		wg.Done()
		clientWg.Done()
	}()

	conn, err := l.Accept()
	require.NoError(t, err)
	require.NotNil(t, conn)
	defer func() { require.NoError(t, conn.Close()) }()
}

func testIsCloseClient(t *testing.T, wg *sync.WaitGroup, clientWg *sync.WaitGroup, l net.Listener) {
	t.Helper()
	defer func() {
		wg.Done()
		clientWg.Done()
	}()

	conn, err := l.Accept()
	require.NoError(t, err)
	require.NotNil(t, conn)
	defer func() { require.NoError(t, conn.Close()) }()

	request := make([]byte, 1024)
	_, err = conn.Read(request)
	require.NoError(t, err)
}
