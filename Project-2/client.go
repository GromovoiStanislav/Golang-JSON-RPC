package main

import (
    "fmt"
    "net/rpc/jsonrpc"
)

func main() {
    client, err := jsonrpc.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    var reply string
    err = client.Call("MyService.Echo", "Привет, мир!", &reply)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Ответ сервера:", reply)
}