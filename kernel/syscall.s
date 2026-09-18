#include "textflag.h"
// NOSPLIT -> evita que exista checagem de crescimento de memória (interrompe codigo no meio) para essa função, interrupção quebraria essa função
// TEXT -> indica início função
// SB -> StaticBase, pseudo-registrador que representa o endereço da função
// DETALHAMENTO SYSCALL:
// 1) Guarda endereço da prox instrução em RCX
// 2) Muda o nível de privilégio
// 3) Desvia instrução para endereço fixo e lê AX e ve qual o código para identificar rotina
// 4) executa lógica do número em AX
// 5) Desfaz o privilégio e volta para próx instrução em RCX
TEXT ·syscall(SB), NOSPLIT, $0-64 // 7 argumentos de 8 bytes + 1 retorno de 1 byte (8*8=64)
    MOVQ num+0(FP), AX     // lê o argumento "num" (offset 0) pro registrador AX
    MOVQ rdi+8(FP), DI     // lê o argumento "rdi" (offset 8) pro registrador DI
    MOVQ rsi+16(FP), SI    // lê o argumento "rsi" (offset 16) pro registrador SI
    MOVQ rdx+24(FP), DX    // lê o argumento "rdx" (offset 24) pro registrador DX
    MOVQ r10+32(FP), R10   // lê o argumento "r10" (offset 32) pro registrador R10
    MOVQ r8+40(FP), R8     // lê o argumento "r8" (offset 40) pro registrador R8
    MOVQ r9+48(FP), R9     // lê o argumento "r9" (offset 48) pro registrador R9
    SYSCALL                // Instrução da CPU
    MOVQ AX, ret+56(FP)    // Syscall guarda file descriptor no AX, move para offset 56 (posição de retorno, onde acaba os argumentos)
    RET // precisa acabar em newline
