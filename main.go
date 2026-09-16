package main

import (
	"fmt"
	"unsafe"

	"server_framework/kernel"
)

func main() {
	mensagem := []byte("oi, isso saiu via MEU proprio syscall\n")

	// write(fd=1, buffer, tamanho)
	// Precisamos de unsafe.Pointer pra passar o endereço real do buffer Go
	// pro assembly -- isso é o "preço" de sair da proteção normal do Go.
	ret := kernel.RawSyscall3(
		kernel.SYS_WRITE,
		1, // fd = stdout
		uintptr(unsafe.Pointer(&mensagem[0])),
		uintptr(len(mensagem)),
	)

	fmt.Println("retorno cru da syscall:", ret)
}