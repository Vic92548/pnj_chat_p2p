package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "192.168.0.25"
	CONN_PORT = "2019"
	CONN_TYPE = "tcp"
)

func main() {

	users := make(map[net.Addr]net.Conn)
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		users[conn.RemoteAddr()] = conn
		// Handle connections in a new goroutine.
		go handleRequest(conn,users)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, users map[net.Addr]net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Print(reqLen)
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.


	// Close the connection when you're done with it.

	size := int(buf[0])
	fmt.Print(size)
	fmt.Print("Size : " + string(int(size)) + "\n")

	msg_buf := make([]byte, size)
	// Read the incoming connection into the buffer.
	length, error := conn.Read(msg_buf)
	if error != nil {
		fmt.Print(length)
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.

	fmt.Print(string(msg_buf) + "\n")
	if(msg_buf[0] == 109){
		msg_buf[0] = 32
		handleMessage(msg_buf,users)
	}

	//conn.Close()
}

func handleMessage(message []byte,users map[net.Addr]net.Conn){
	for k, v := range users {
		v.Write(message)
		print("send to " + k.String() + "\n")
	}
}