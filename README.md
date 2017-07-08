# TcpFileTransfer
A simple TCP file transfer program between a client and a server.

tcp_server.go (SERVER)
A Go TCP server wich receives files from clients.
This program will run a TCP server at localhost:7005 and localhost:7006.
You can connect to this server using the tcp_client.go program.

tcp_client.go (CLIENT)
A Go TCP client wich sends files to a server.
You can connect to that server with default parameters like:
  $ go run tcp_client.go
assuming the server is already running (see above).
Default parameters:
  IP: 127.0.0.1.
  PORT1: 7005.
  PORT2: 7006.
And, of course, you can connect to the server using its IP
address and ports like:
  $ go run tcp_client.go IP-ADDRESS PORT1 PORT2
assuming your firewall allows it.
To send a file:
  $ send 'filename'
    
    
