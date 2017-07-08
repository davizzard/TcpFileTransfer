package main

/*
	tcp_server.go -- A Go TCP server wich receives files from clients.
	This program will run a TCP server at localhost:7005 and localhost:7006.
	You can connect to this server using the tcp_client.go program.
*/

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const BUFFER_SIZE = 1024
const REQUEST_PORT = "7005"
const TRANSFER_FILE_PORT = "7006"

func main() {

	// We will listen from two ports.
	// In the first connection we'll receive the client request type, metadata from the file and
	// other information.
	// The second one is dedicated to file transfer.
	//
	server, err := net.Listen("tcp", "localhost:"+REQUEST_PORT)
	if err != nil {
		fmt.Println("There was an error while starting the server: " + err.Error())
		return
	}

	server2, err := net.Listen("tcp", "localhost:"+TRANSFER_FILE_PORT)
	if err != nil {
		fmt.Println("There was an error while starting the server: " + err.Error())
		return
	}

	fmt.Println("Listening on ports: "+REQUEST_PORT+" and "+TRANSFER_FILE_PORT)

	// We accept incoming connections
	//
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("There was an error with the connection: " + err.Error())
			return
		}
		fmt.Println("A client has connected.")

		connection2, err := server2.Accept()
		if err != nil {
			fmt.Println("There was an error with the connection: " + err.Error())
			return
		}
		// We handle the connection, on it's own thread, per connection
		go connectionHandler(connection, connection2)
	}

}

// connectionHandler handles client request and executes corresponding function
//
func connectionHandler(connection net.Conn, connection2 net.Conn) {
	buffer := make([]byte, BUFFER_SIZE)

	_, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("There is an error reading from connection: ", err.Error())
		return
	}

	arrayOfCommands := strings.Split(string(buffer), " ")

	fmt.Println("COMMAND: " + arrayOfCommands[0])
	fmt.Println("FILENAME: " + arrayOfCommands[1])

	if arrayOfCommands[0] == "send" {
		fmt.Println("Getting file " + arrayOfCommands[1] + " from client...")
		getFileFromClient(arrayOfCommands[1], connection, connection2)
	} else {
		_, err = connection.Write([]byte("Bad command."))
	}

}

// getFileFromClient receives file data from client, saves it into
// a new file with the same name as the original and closes both connections
func getFileFromClient(fileName string, connection net.Conn, connection2 net.Conn) {
	file, err := os.Create(strings.TrimSpace(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	defer connection.Close()
	defer connection2.Close()

	n, err := io.Copy(file, connection2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	file.Sync()
	fmt.Println(n, "Bytes recieved.\n")

	return

}
