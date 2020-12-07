extern printf

section .data
    message db "RAX: %d", 0AH, 0H

section .text
global main
main:

    push rbp

    xor rax, rax
    mov rbx, 100
    soma:
    add rax, rbx
    dec rbx
    cmp rbx, 0
    jne soma
    
    mov rdi, message
    mov rsi, rax
    xor rax, rax
    call printf
    
    pop rbp
    
    xor rax, rax
    ret    