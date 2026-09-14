package main

import (
	"syscall"
)

func main() {
	mensagem := []byte("oi, isso saiu direto via syscall.Write\n")

	// syscall.Write(fd, buffer) -> (n int, err error)
	// fd = 1 é o stdout, por convenção do POSIX (0=stdin, 1=stdout, 2=stderr)
	// isso NÃO passa por fmt.Println nem por buffer do runtime -- é a
	// chamada de sistema "write" sendo disparada diretamente
	n, err := syscall.Write(1, mensagem)
	if err != nil {
		panic(err)
	}

	println("bytes escritos:", n)
}