package kernel

// RawSyscall3 está implementada em syscall_amd64.s
// Aqui só declaramos a assinatura -- sem corpo, porque o corpo é assembly.
func RawSyscall3(num, a1, a2, a3 uintptr) uintptr

// Números de syscall do Linux x86-64 (vêm da tabela oficial do kernel,
// arquivo arch/x86/entry/syscalls/syscall_64.tbl)
const (
	SYS_WRITE = 1
)