extern printf

section .data
    format db "a = %d", 0AH, 0H

section .text
global main
main:

    push rbp

    mov rbx, 1000
    mov rcx, 11
    
    divisao:
    mov rax, rbx
    mov rdx, 0
    div rcx
    cmp rdx, 5
    jne nprint
    
    push rax
    push rcx
    
    mov rdi, format
    mov rsi, rbx
    xor rax, rax
    call printf
    
    pop rcx
    pop rax
    
    nprint:
    inc rbx
    cmp rbx, 2000
    jne divisao
    
    pop rbp    
    
    xor rax, rax
    ret