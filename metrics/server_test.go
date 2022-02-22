package main

import (
	"fmt"
	"testing"
	"time"
)

/*
Tests the server startup.

NOTE:
When running or debugging this test, you should have a test_config.env file in your development folder.
The test_config.env should have the uri environment variable defined with a working MongoDB URI. (place your real URI, that file is in the .gitignore)
*/
func TestStartServer(t *testing.T) {
	fmt.Println("server started")

	go (func() {
		StartServer() // running a goroutine which will terminate at the end of the test, stopping all subprocesses, hence terminating the server
	})()

	time.Sleep(time.Second * 2)

	fmt.Println("Killing server")
}
