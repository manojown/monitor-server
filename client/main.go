package main

import (
	"encoding/json"
	"fmt"
	"net"
	"serverInfo/client/services"
)

type Client struct {
	socket net.Conn
	data   chan []byte
}

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("RECEIVED: " + string(message))
		}
	}
}

func startClientMode() {
	fmt.Println("Starting client...")
	connection, error := net.Dial("tcp", "ec2-52-91-69-245.compute-1.amazonaws.com:12345")
	if error != nil {
		fmt.Println(error)
	}
	client := &Client{socket: connection}
	go client.receive()
	v := OsUtility.GetMem()
	OsUtility.GetDisk()

	// fmt.Printf("Asdasdasdasd data is", disk)
	d, _ := json.Marshal(v)
	connection.Write([]byte(d))

}

func main() {

	startClientMode()

}
