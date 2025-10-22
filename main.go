package main

import (
	"fmt"
	"os"
	"pkg-registry/server"
)

func main() {
	err := server.StartServer()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
}
