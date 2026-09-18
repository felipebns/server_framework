package kernel

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

func write()

//precisa retornar uintpr pq o file descriptor pode ser um ponteiro, desse modo consegue armazenar tanto inteiro quanto ponteiros
func Socket(domain int, typ int, protocol int) uintptr { 
	//num -> SYS_SOCKET (código de socket)
	//rdi -> domain (padrão x86-64)
	//rsi -> typ (padrão x86-64)
	//rdx -> protocol (padrão x86-64)
	return syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(protocol), 0, 0, 0) 
}

func bind()

func listen()

func accpet()

func read()