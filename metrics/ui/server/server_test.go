package main

import (
	"fmt"
	"testing"
	"time"
)

func TestStartServer(t *testing.T) {
	fmt.Println("server started")

	go (func() {
		StartServer()
	})()

	time.Sleep(time.Second * 2)

	fmt.Println("Killing server")
}
