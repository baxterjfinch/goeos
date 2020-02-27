package main

import (
    "fmt"
    eosapi "github.com/hemlokc/goeos/src/rpc"
)

func main() {
	fmt.Println("Hello, world.")

    connection := rpc.RPCDetails {
        Endpoint: "http://localhost:8888",
    }

    connection.FetchConnection()
}
