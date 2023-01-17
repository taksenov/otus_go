// Package hw11_telnet_client -- OTUS HW11 Telnet Client.
package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "connection timeout")
}

func main() {
	flag.Parse()

	flagArgs := flag.Args()
	if len(flagArgs) != 2 {
		log.Fatal("usage: 'go-telnet [--timeout=10s] host port'")
	}
	address := net.JoinHostPort(flagArgs[0], flagArgs[1])

	c := NewTelnetClient(address, timeout, ioutil.NopCloser(os.Stdin), os.Stdout)
	log.Printf("...Connected to %s\n", address)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	// Graceful Shutdown
	go func() {
		<-gracefulShutdown
		cancel()
	}()

	// Send
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := c.Send(); err != nil {
					c.Close()
					log.Println(err)
					cancel()

					return
				}
			}
		}
	}()

	// Receive
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := c.Receive(); err != nil {
					c.Close()
					log.Println(err)
					cancel()

					return
				}
			}
		}
	}()

	wg.Wait()
}
