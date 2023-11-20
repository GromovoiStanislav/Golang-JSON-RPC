package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type MyService struct{}

func (s *MyService) Echo(args string, reply *string) error {
    *reply = args
    return nil
}

func main() {
   
	rpc.Register(new(MyService))

    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }

	fmt.Println("Сервер запущен на порту 8080")

    for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

   
