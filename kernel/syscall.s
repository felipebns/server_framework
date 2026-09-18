#include "textflag.h"
// NOSPLIT -> evita que exista checagem de crescimento de memória (interrompe codigo no meio) para essa função, interrupção quebraria essa função
// TEXT -> indica início função
// SB -> StaticBase, pseudo-registrador que representa o endereço da função
TEXT ·syscall(SB), NOSPLIT, $0-64 // 7 argumentos de 8 bytes + 1 retorno de 1 byte (8*8=64)
    MOVQ num+0(FP), AX     // lê o argumento "num" (offset 0) pro registrador AX
    MOVQ AX, ret+56(FP)  // escreve o resultado no offset 56 (posição de retorno, onde acaba os argumentos)
    RET // precisa acabar em newline
