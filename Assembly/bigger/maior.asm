 
SYS_EXIT  equ 1
RET_EXIT  equ 5
SYS_READ  equ 3
SYS_WRITE equ 4
STD_IN    equ 0
STD_OUT   equ 1
MAX_IN    equ 10
 
section .data
    fmsg db "Digite dois numeros: (int)", 0xA
    flen equ $-fmsg
    
    mai db "O primeiro numero e maior", 0xA
    glen equ $-mai
    
    men db "O primeiro numero e menor", 0xA
    llen equ $-men
    
    ig db "Os numeros sao iguais", 0xA
    elen equ $-ig
    
section .bss

    num1 resb 4
    num2 resb 4
    
section .text

    global _start
    
    _start:

    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, fmsg
    mov edx, flen
    int 0x80
    
    mov eax, SYS_READ
    mov ebx, STD_IN
    mov ecx, num1
    mov edx, MAX_IN
    int 0x80
    
    mov eax, SYS_READ
    mov ebx, STD_IN
    mov ecx, num2
    mov edx, MAX_IN
    int 0x80
    
    mov eax, SYS_WRITE
    mov ebx, STD_OUT
    mov ecx, [num1]
    cmp ecx, [num2]
    jl menor
    jg maior
    je igual
    
    igual:
    mov ecx, ig
    mov edx, elen
    int 0x80
    jmp exit
    
    menor:
    mov ecx, men 
    mov edx, llen
    int 0x80
    jmp exit
    
    maior:
    mov ecx, mai 
    mov edx, glen
    int 0x80
    jmp exit
    exit:
    
    mov eax, SYS_EXIT
    mov ebx, RET_EXIT
    int 0x80
    
    
    
    
    
    
