package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	args := Args{3, 4}
	var result int
	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Println("Error calling remote procedure:", err)
		return
	}

	fmt.Printf("Result: %d\n", result)
}
