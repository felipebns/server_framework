package erro_handler

import "fmt"

func HandleSys(erro int64) int {
	if erro < 0 {
		if erro == -13 {
			fmt.Println("Erro: permissão negada (EACCES) — provavelmente tentando usar porta menor que 1024 sem privilégio de root")
		} else if erro == -98 {
			fmt.Println("Erro: endereço já em uso (EADDRINUSE) — outro processo já está associado a esse IP/porta")
		} else if erro == -99 {
			fmt.Println("Erro: endereço indisponível (EADDRNOTAVAIL) — esse IP não pertence a nenhuma interface local")
		} else if erro == -9 {
			fmt.Println("Erro: file descriptor inválido (EBADF) — o fd passado não corresponde a um socket aberto")
		} else if erro == -97 {
			fmt.Println("Erro: família de endereço não suportada (EAFNOSUPPORT) — domain passado incorretamente")
		} else {
			fmt.Printf("Erro desconhecido, código: %d\n", erro)
		}
		return -1
	}
	return 0
}