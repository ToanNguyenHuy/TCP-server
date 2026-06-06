package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func TestS(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:1729")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	status, err := reader.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(status) != "HTTP/1.1 200 OK" {
		t.Fatalf("unexpected status: %q", status)
	}
}
