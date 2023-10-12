
        section .data
buffer:
        db 100                         ; Reserva espacio para almacenar los datos leídos

        section .text
        global _start

_start:
; Configura los argumentos para la llamada al sistema read
        mov rax, 0                     ; Número de llamada al sistema para read
        mov rdi, 0                     ; Descriptor de archivo 0 (stdin para entrada estándar)
        mov rsi, buffer                ; Dirección de memoria del búfer donde se almacenarán los datos leídos
        mov rdx, 100                   ; Cantidad de bytes a leer (tamaño del búfer)

; Llama al sistema read
        syscall

; En este punto, los datos leídos se encuentran en el búfer 'buffer'

; Tu código para trabajar con los datos leídos aquí

; Salir del programa
        mov rax, 60                    ; Número de llamada al sistema para exit
        xor rdi, rdi                   ; Código de salida 0
        syscall

