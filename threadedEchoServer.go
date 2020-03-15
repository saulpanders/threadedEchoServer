/*
	@saulpanders
	8/25/18
	threadedEchoServer.go --- echo's client response back w/ threads to handle multiple client connections

	to try it, go run/build then use ncat
		ncat 127.0.0.1 1200

*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	service := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//run client handling as subroutine (goroutine)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//close connection on exit
	defer conn.Close()

	var buf [512]byte
	for {

		//read up to 512 bytes at once
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		//write n bytes read back to client
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
