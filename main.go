package main

import (
	"log"
	"net"
	"time"
)

// Block call
func do(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World! s\r\n"))
	conn.Close()
}

func main() {
	log.Println("Start Server...")
	// When we start server, we want to listen to a particular port
	// This process reserve a port
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, er := listener.Accept()
		if er != nil {
			log.Fatal(er)
		}
		do(conn)
	}

}
