 
 
 SYS_EXIT  equ 1
 RET_EXIT  equ 5
 SYS_READ  equ 3
 SYS_WRITE equ 4
 STD_IN    equ 0
 STD_OUT   equ 1
 MAX_IN    equ 10
 
 segment .data
    msg db "Digite seu nome: ", 0xA, 0xD
    len equ $-msg
    gre db "ola ", 0xE    
    len2 equ $-gre
    
segment .bss
    nome resb 2
    
segment .text

    global _start
    
_start:

    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, msg
    mov edx, len
    int 0x80
    
    ; data input
    
    mov eax, SYS_READ
    mov edx, STD_IN
    mov ecx, nome
    mov edx, MAX_IN
    int 0x80
    
    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, gre
    mov edx, len2
    int 0x80
    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, nome
    mov edx, 10
    int 0x80
    
exit:

    mov eax, SYS_EXIT
    mov ebx, RET_EXIT
    int 0x80
    
    
    
    
    
    
    
    
    
    
    
    
    
