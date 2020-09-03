package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	// Constant defination.
	ADDRESS = "127.0.0.1:8081"
)

func main() {

	//Connect to the server
	fmt.Println("Connecting to server.")
	conn, err := net.Dial("tcp", ADDRESS)
	if err != nil {
		fmt.Println("Error encountered while making connection is - ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to server.")

	for {
		//Ask specified user for a message.
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter a message: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error encountered while reading the message is - ", err.Error())
			fmt.Println("Exiting client.")
			os.Exit(1)
		}

		//Send the mesasge to server.
		conn.Write([]byte(msg))

		//Wait for acknowledgement from server.
		acknowledge, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error occured while receiving acknowledgement from server.")
			panic(err)
		}
		//Validate the acknowledgement.
		if acknowledge == "Message received\n" {
			fmt.Println("Server received message successfully.")
		} else {
			fmt.Println("Message not sent. Send again!")
		}
	}
}
