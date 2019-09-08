package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "192.168.0.25:2019")
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
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nMessage from server: "+message)
	}
}
