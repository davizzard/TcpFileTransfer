package main

/*
	tcp_client.go -- A Go TCP client wich sends files to a server.
	You can connect to that server with default parameters like:
		> go run tcp_client.go
	(assuming the server is already running -> tcp_server.go).
	Default parameters
		--> IP: 127.0.0.1.
		--> PORT1: 7005.
		--> PORT2: 7006.
	And, of course, you can connect to the server using its IP
	address and ports like:
		> go run tcp_client.go IP-ADDRESS PORT1 PORT2
	assuming your firewall allows it.

	To send a file:
		> send <filename>
*/


import (
	"fmt"
	"net"
	"os"
	"io"
	"bufio"
	"strings"
)


func main() {

	//We get ports and ip address to dial.
	//
	var ip string
	var r_port string
	var f_port string
	if len(os.Args) != 4 {
		fmt.Println("\nDefault Server ip: 127.0.0.1.")
		fmt.Println("Default ports: 7005 and 7006.")
		fmt.Println("Example of use:\n\t> tcp_client.go <server-ip-address> <server-port1> <server-port2>")
		ip = "127.0.0.1"
		r_port = "7005"
		f_port = "7006"
	} else {
		ip = os.Args[1]
		r_port = os.Args[2]
		f_port = os.Args[3]
	}

	reader := bufio.NewReader(os.Stdin)

	// Loop endlessly.
	//
	for {
		// We open new connections for each client request.
		connection, err := net.Dial("tcp", ip+":"+r_port)
		if err != nil {
			fmt.Println("There is an error reading from connection: ", err.Error())
			return
		}

		connection2, err := net.Dial("tcp", ip+":"+f_port)
		if err != nil {
			fmt.Println("There is an error reading from connection: ", err.Error())
			return
		}
		fmt.Print("Please enter 'send <filename>' to transfer a file to the server.\n")

		inputFromUser, _ := reader.ReadString('\n')
		arrayOfCommands := strings.Split(inputFromUser, " ")

		// We process user's input and call the corresponding function.
		if arrayOfCommands[0] == "send" {
			sendFileToServer(arrayOfCommands[1], connection, connection2)
		} else {
			fmt.Println("Bad Input.")
		}
	}

}

// sendFileToServer transfers input file data to server via TCP and closes both connections.
func sendFileToServer(fileName string, connection net.Conn, connection2 net.Conn) {
	// For read access.
	file, err := os.Open(strings.TrimSpace(fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	defer connection.Close()
	defer connection2.Close()

	connection.Write([]byte("send " + fileName + " "))

	fmt.Println("Sending file " + fileName + " to server...")
	n, err := io.Copy(connection2, file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n, "Bytes sent")

	return

}