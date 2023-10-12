
        section .data
hello:
        db 'Hola, mundo!', 10, 0       ; Definir una cadena de caracteres terminada en null

        section .text
        global _start

_start:
; Configurar los registros para la llamada al sistema write
        mov rax, 1                     ; Número de llamada al sistema para write
        mov rdi, 1                     ; Descriptor de archivo para stdout (salida estándar)
        mov rsi, hello                 ; Dirección del mensaje
        mov rdx, 13                    ; Longitud del mensaje

; Llamar al sistema write
        syscall

; Salir del programa
        mov rax, 60                    ; Número de llamada al sistema para exit
        xor rdi, rdi                   ; Código de salida 0
        syscall

