### Máquina e Linguagem:
Considera-se que o servidor vai rodar em uma máquina Linux e arquitetura x86-64

Para tal, a linguagem Go é utilizada.

Fatores decisivos na escolha:
- Go é uma linguagem com compilador que gera código de máquina nativo (instruções que a CPU executa diretamente)
- A linguagem expõe formas de alterar o valor de registradores e disparar instruções syscall (alterações em nivel de kernel)

### Assembly:
Plan 9 assembly, não é assembly x86-64 puro, é uma camada de abstração por cima das instruções reais da CPU, existem pseudo-registradores que facilitam o acesso a variáveis guardadas na memória. O próprio assembler GO que traduz as abstrações em endereçamento real quando vai montar o binário. 

- Compilador GO empilha os argumentos na memória, na ordem que aparecem na assinatura, ocupando o número de bytes correspondente ao seu tipo, FP marca onde a área começa, puxa os argumentos a partir dele como em uma stack (0 - 8 - 16 - etc ....)