
; || CODE.

        section .text
        global _start

_start:
        mov rax, 1                     ; Syscall write.
        mov rdi, 1                     ; File descriptor (stdout).
        mov rsi, prompt
        mov rdx, prompt_length
        syscall

        mov rax, 0                     ; Syscall read.
        mov rdi, 0                     ; File descriptor (stdin).
        mov rsi, name                  ; Save in name.
        mov rdx, name_length           ; Total of bytes to read.
        syscall

        mov r12, rax                   ; Save return value of syscall read.
        sub r12, 1                     ; Exclude newline (last character).

        mov rax, 1                     ; Syscall write.
        mov rdi, 1                     ; Stdout.
        mov rsi, text1
        mov rdx, text1_length
        syscall

        mov rax, 1
        mov rdi, 1, 
        mov rsi, name
        mov rdx, r12                   ; Check note of bottom of syscall read.
        syscall

        mov rax, 1
        mov rdi, 1
        mov rsi, text2
        mov rdx, text2_length
        syscall

        mov rax, 60                    ; Syscall exit.
        mov rdi, 0                     ; Exit code 0.
        syscall

; || DATA.

        section .data

prompt:
        db "Enter your name: "
        prompt_length equ $ - prompt

text1:
        db 10, "Hello!, "
        text1_length equ $ - text1

text2:
        db ". This is assembler!...", 10
        text2_length equ $ - text2

name:
        db 100                         ; Reservar 100 bytes, para el nombre del usuario.
        name_length equ 100

