package main

import (
	"bufio"
	"net"
	"os"
	"time"
)
import "fmt"

func main() {


	// connect to this socket
	conn, _ := net.Dial("tcp", "192.168.0.11:2019")

	go socket_reader(conn)


	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nText to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		l := byte(len(text))
		msg := []byte(" " + "m" + text)
		msg[0] = l
		conn.Write(msg)
	}
}

func socket_reader(conn net.Conn) {
	var name string = "jean michel"
	var time_sent string = time.Now().Format("2006-01-02 15:04:05")
	for {
		// listen for reply
		buf := make([]byte, 1)
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Print(reqLen)
			fmt.Println("Error reading:", err.Error())
		}

		size := int(buf[0])
		print(size)
		msg_buf := make([]byte, size)

		reqLen2, err2 := conn.Read(msg_buf)
		if err2 != nil {
			fmt.Print(reqLen2)
			fmt.Println("Error reading:", err2.Error())
		}
		//fmt.Print("\nMessage from server: "+ string(msg_buf))

		if(msg_buf[0] == 109){
			go handle_message(string(msg_buf[1:]),name,time_sent)
		}
		if (msg_buf[0] == 105) {
			time_sent = string(msg_buf[1:])
		}
		if (msg_buf[0] == 110) {
			name = string(msg_buf[1:])
		}
	}
}

func handle_message(message string, name string, time string){
	fmt.Print(time + " - " + name + " : " + message)
}