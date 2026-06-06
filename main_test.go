package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func TestS(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:1729")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(res) != "Hello, World! s" {
		t.Fatalf("Expected Hello")
	}
}
