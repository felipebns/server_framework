package main

import (
	"fmt"
	"server_framework/kernel"
)

func main() {
	file_descriptor := kernel.Socket(2, 1, 0)
	porta := 8000
	ip := "0.0.0.0"
	r := kernel.Bind(file_descriptor, porta, ip)
	fmt.Println("File descriptor:", file_descriptor)
	fmt.Println("Bind result:", r)
}