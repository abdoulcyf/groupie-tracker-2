package main

import (
	"fmt"
	"main/internal/servers"
)

func main() {
	errServer := servers.WebServer()
	if errServer != nil {
		fmt.Println(errServer)
	} else {
		fmt.Println("Server is running on port 8080")
	}
}
