package main

// * Example of usage * //

import (
	"fmt"
	"log"
	"main/goports"
	"os"
)

func main() {

	// TODO: Create Adapter for CLI-arguments to Inject into goports.Scan()

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run scanner.go <IP> <Port/PortRange>")
		fmt.Println("Example: go run scanner.go 127.0.0.1 80")
		fmt.Println("Example: go run scanner.go 127.0.0.1 1-20000")
		return
	}

	ip, err := goports.ParseIP(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	port_begin, port_end, err := goports.ParsePorts(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	goports.Scan(ip, port_begin, port_end)
}
