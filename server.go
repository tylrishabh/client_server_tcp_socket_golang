package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	//Server will be hosted on this port.
	PORT = ":8081"

	//Ackowledgement message to be sent to client.
	ACKNOWLEDGE = "Message received\n"
)

func main() {

	//Enables the server to accept connections on port.
	fmt.Println("Listening on port", PORT)
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error encountered while making connection is - ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Waiting to connect.")
	//Accept connection on port.
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error occured while accepting is - ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Connection accepted successfully.")
	defer conn.Close()

	for {
		//Blocking call that waits until a message is received.
		reader, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error occured while reading the message is - ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Ready to receive message\n")
		fmt.Println("Message from client is : ", string(reader))

		//Acknowledge the client that its message is received successfully.
		conn.Write([]byte(ACKNOWLEDGE))
	}
}
