### Pacote próprio para comunicação com o kernel

- Criei meu próprio pacote para realizar a comunicação com o karnel através de instruções que serão enviadas até a CPU via código compilado.

- O objetivo (geral) é um programa que roda em um ring de privilégio menor solicite um nível maior, a fim de executar uma operação. Mecanismo chamado de *Syscall*

- Toda informação de uma syscall precisa ser posicionada em registradores especificos antes da instrução ser executada

* RAX: número identificador da syscall desejada 

    write = 1

    socket = 41

    bind = 49

    listen = 50

    accept = 43

    read = 0

* RDI, RSI, RDX, R10, R8 e R9 carregam atpe seis argumentos na chamada

* Retorno: RAX (negativo = erro, positivo = erro padronizado)


### O que exatamente essas operações fazem:

- Socket: Solicita ao kernel a criação de um novo endpoint de comunicação, especificando família de endereço e tipo (TCP, UDP, etc). O kernel aloca uma estrutura de dados que representa esse endpoint e devolve um identificador numérico para ele - file descriptor.

- Bind: Associa o file descriptor a um endereço IP e porta específicos

- Listen: muda o estado interno daquela entrada de "fechada" para "escutando" e defino o tamanho de uma fila interna que vai acumular conexões recebidas 

- Accept: Retira da fila interna uma conexão já totalmente estabelecida e devolve um novo file descriptor, referente a essa conversa com aquele cliente

- Read/recv: solicita ao kernel os bytes que já chegaram e estão armazenados num buffer de recepção referentes àquele file descriptor de conexão

- Write: entrega ao kernel os bytes que seu programa deseja transmitir

- Close: livera o file descriptor e sinaliza ao kernel que aquela conexão específica pode ser encerrada