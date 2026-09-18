package main

import (
	"fmt"
	"server_framework/kernel"
)

func main() {
	file_descriptor := kernel.Socket(0, 0, 0)
	fmt.Println(file_descriptor)
}