package main

import (
	"fmt"
	"monitor-server/client/services"
	"net"
	"os"

	"github.com/subosito/gotenv"
)

type Client struct {
	socket net.Conn
	data   chan []byte
}

// type DetailedManager SharedModel.DetailedManager

func init() {
	gotenv.Load()
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
	MasterData := make(chan []byte)
	connection, error := net.Dial("tcp", os.Getenv("SERVER_URL"))
	if error != nil {
		fmt.Println(error)
	}
	// client := &Client{socket: connection}

	go OsUtility.GetMem(MasterData)
	go OsUtility.GetDisk(MasterData)
	go OsUtility.GetCpu(MasterData)
	go OsUtility.GetDocker(MasterData)

	for {
		select {
		case state := <-MasterData:
			connection.Write([]byte(state))

		}
	}
}

func main() {

	startClientMode()

}
