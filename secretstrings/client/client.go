package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"uk.ac.bris.cs/distributed1/chat/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect as server")
	flag.Parse()
	fmt.Println("Server:", server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded:" + response.Message)
	// Connect to the RPC server and send the request(s)

}
