package main

import (
	"fmt"
	"net"
)

func main() {
	ifs, err := net.Interfaces()
	fmt.Printf("%v, %v", ifs, err)
}
