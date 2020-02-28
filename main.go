package main

import (
    "fmt"
    eosapi "github.com/hemlokc/goeos/src/rpc"
)

func main() {
	fmt.Println("Hello, world.")

    api := eosapi.New("http://localhost:8888")

    fmt.Println(api)
}
