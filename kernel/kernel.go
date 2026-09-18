package kernel

import "server_framework/utils"
import "encoding/binary"
import "strings"
import "strconv"
import "unsafe"
import "fmt"

// Declarando assinaturas de funções externas (assembly) que implementam syscalls.
// Registador RAX vai receber o identifier
// Retorno é um file descriptor (int)
// syscall recebe identifier e é chamada de write/socket, etc, é código repetitivo
func syscall(num uintptr, rdi uintptr, rsi uintptr, rdx uintptr, r10 uintptr, r8 uintptr, r9 uintptr) uintptr

// Números de syscall do Linux x86-64 (vêm da tabela oficial do kernel,
// arquivo arch/x86/entry/syscalls/syscall_64.tbl)
const (
	SYS_WRITE = 1
	SYS_SOCKET = 41
	SYS_BIND = 49
	SYS_LISTEN = 50
	SYS_ACCEPT = 43
	SYS_READ = 0
)

// Struct para definir IP e porta
type sock_addr_in struct {
	// sin_family é litle endian (campo interno, fora da convenção)
	// port e addr são big endians (convenção de quando criaram protocolo de rede)
	sin_family [2]byte // família/domínio de endereço - AF_NET
	sin_port [2]byte // porta do servidor
	sin_addr [4]byte // endereço IP do servidor
	sin_zero [8]byte // padding (cheio de zeros)
}

//precisa retornar uintpr pq o file descriptor pode ser um ponteiro, desse modo consegue armazenar tanto inteiro quanto ponteiros
func Socket(domain int, typ int, protocol int) uintptr { 
	//num -> SYS_SOCKET (código de socket)
	//rdi -> domain (padrão x86-64) (família de endereço, avisa para o kernel o tipo de dado que vai entrar, AF_INET = 2)
	//rsi -> typ (padrão x86-64) (tipo conexão, SOCK_STREAM = 1)
	//rdx -> protocol (padrão x86-64) (protocolo da conexão, TCP = 0)
	r := syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(protocol), 0, 0, 0) 

	// próprio kernel vai rejeitar se os parametros vierem errado
	erro := int64(r)
	erro_handler.HandleSys(erro)
	return r
}

// porta e endereço já chegam decididos
func Bind(fd uintptr, porta int, endereco string) int {
	// fd é file descriptor do socket
	// sock_addr_in é um struct em que definiremos o endereço do servidor
	// len_addr é o tamanho desse struct (sempre 16 bytes)
	if porta < 0 || porta > 65535 {
		fmt.Println("Porta inválida. Deve estar entre 0 e 65535.")
		return -1
	}

	uint_porta := uint16(porta)
	buffer_porta := [2]byte{0,0}
	binary.BigEndian.PutUint16(buffer_porta[:], uint_porta) //espera um slice

	r := strings.Split(endereco, ".")
	if len(r) != 4 {
		fmt.Println("Endereço inválido, precisa ter 4 seções")
		return -1
	}

	var ip_completo []byte
	for _, section_ip := range r {
		if len(section_ip) > 3 {
			fmt.Println("Endereço inválido")
			return -1
		}
		valor, err := strconv.Atoi(section_ip) // converte string para int
		if valor < 0 || valor > 255 {
			fmt.Println("Cada seção do IP deve estar entre 0 e 255")
			return -1
		}

		if err != nil {
			fmt.Println("Endereço contém caractere inválido")
			return -1
		}

		byte_ip := byte(valor)
		ip_completo = append(ip_completo, byte_ip)
	}

	sock_completo := sock_addr_in{
		sin_family: [2]byte{2,0}, //litle endian (sempre AF_INET = 2)
		sin_port: buffer_porta, // big
		sin_addr: [4]byte{ip_completo[0], ip_completo[1], ip_completo[2], ip_completo[3]}, // big por natureza
		sin_zero: [8]byte{0,0,0,0,0,0,0,0},
	}
	
	// unsafe pointer pq é necessário converter a struct em uintptr (formato que o assembly recebe, também é o formato do ponteiro)
	// 16 pq é o tamanho do sock_completo em bytes
	r_sys := syscall(SYS_BIND, fd, uintptr(unsafe.Pointer(&sock_completo)), 16, 0, 0, 0) 
	erro := int64(r_sys)
	handled_error := erro_handler.HandleSys(erro)
	return handled_error
}

func Listen()

func Accpet()

func Read()

func Write()

func Close()