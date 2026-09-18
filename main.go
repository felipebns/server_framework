package main

import (
	"fmt"
	"server_framework/kernel"
)

func main() {
	teste := kernel.Socket(0, 0, 0)
	fmt.Println(teste)
}