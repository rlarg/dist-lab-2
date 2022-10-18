package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed1/chat/stubs"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error: exiting")
		return
	}
}

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect as server")
	flag.Parse()
	fmt.Println("Server: ", server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	//request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)
	file, err := os.Open("wordlist")
	check(err)
	rd := bufio.NewReader(file)
	for {
		line, _ := rd.ReadString('\n')
		request := stubs.Request{Message: line}
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded: " + response.Message)
	}
	file.Close()
	// Connect to the RPC server and send the request(s)
}
