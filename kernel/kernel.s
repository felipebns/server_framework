#include "textflag.h"

// func RawSyscall3(num, a1, a2, a3 uintptr) uintptr
// Monta os registradores na convenção do Linux x86-64 e dispara SYSCALL.
TEXT ·RawSyscall3(SB), NOSPLIT, $0-40
    MOVQ num+0(FP), AX   // rax = número da syscall
    MOVQ a1+8(FP), DI    // rdi = primeiro argumento
    MOVQ a2+16(FP), SI   // rsi = segundo argumento
    MOVQ a3+24(FP), DX   // rdx = terceiro argumento
    SYSCALL              // dispara a instrução real de CPU
    MOVQ AX, ret+32(FP)  // resultado (ou erro negativo) volta em rax
    RET
// precisa de um \n no final!
