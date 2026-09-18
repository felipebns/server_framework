### Máquina e Linguagem:
Considera-se que o servidor vai rodar em uma máquina Linux e arquitetura x86-64

Para tal, a linguagem Go é utilizada.

Fatores decisivos na escolha:
- Go é uma linguagem com compilador que gera código de máquina nativo (instruções que a CPU executa diretamente)
- A linguagem expõe formas de alterar o valor de registradores e disparar instruções syscall (alterações em nivel de kernel)

### Assembly:
