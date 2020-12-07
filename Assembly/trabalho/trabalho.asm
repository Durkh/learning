extern scanf, printf

section .bss
    integer1   resd 4

section .data
    request db "Digite o número: ", 0H
    message db 0AH, "O elemento %d da sequencia é: %d", 0AH, 0H
    input   db "%d", 0
    negerrormsg db 0AH, "Erro: Digite um número positivo.", 0AH, 0H

section .text
global main
main:
    push    rbp             ;alinha o stack
    
    mov     rdi, request    ;mensagem inicial
    xor     rax, rax
    call    printf
    
    mov     rsi, integer1   ;input do numero
    mov     rdi, input
    xor     rax, rax
    call    scanf
    
    mov     rax, [integer1] ;checa se o número é negativo
    mov     rbx, 0x80000000 ;testei várias coisas mas só consegui ->
    test    rax, rbx        ;fazendo o test pra ver se o bit de negativo tava on
    jnz     neg_error
    
    mov     rdi, rax        ;vi que é bom costume usar rdi para parametro
    call    fib              ;chama a função
    
    print:                  ;printa o numero da sequência
    mov     rdx, rax        ;(pensei em colocar em uma variável mas como eu sei ->
    mov     rsi, [integer1] ; que só vou usar rax aqui, deixei logo em rax, deve ->
    mov     rdi, message    ; ser assim a implementação de uma chamada de função no printf)
    call    printf
    
    jmp     finish
    
    neg_error:              ;printa erro de numero negativo
    mov     rdi, negerrormsg
    xor     rax, rax
    call    printf
    
    finish:                 ;e aqui acaba a festa
    
    pop     rbp
    
    xor     rax, rax       ;retorna 0
    ret
    
fib: 
    push    rbp         ;setando o stack frame com 2 variáveis
    mov     rbp, rsp
    push    rbx
    sub     rsp, 8
    mov     qword [rbp-16], rdi
    
    cmp     qword [rbp-16], 0
    jne     case1       ;primeiro if 
    mov     rax, 0
    jmp     end
    
    case1:              
    cmp     qword [rbp-16], 1
    jne     recursao    ;segundo if
    mov     rax, 1
    jmp     end
    
    recursao:           ;recursão e soma
    mov     rax, qword [rbp-16]
    dec     rax
    mov     rdi, rax
    call    fib
    mov     rbx, rax
    mov     rax, qword [rbp-16]
    sub     rax, 2
    mov     rdi, rax
    call    fib
    add     rax, rbx
    
    end:                ;retorno da função
    mov     rbx, qword [rbp-8] 
    leave               ;descobri esse opcode, útil, ele destrói o stack frame
    ret
