
# Assembly Practice

Práctica en lenguaje ensamblador. Aquí añadiré los conceptos que vaya aprendiendo.

> **ADVERTENCIA**: Esto solo aplica para **Linux AMD64 (NASM)**.

## Resumen

### Funciones de ensamblador

Cada llamada al sistema tiene un número asignado, este número se debe enviar al registro **rax**, después de le envían los argumentos en el correspondiente registro y se ejecuta **syscall**.

#### Argumentos de funciones

 - rdi: Primer argumento.
 - rsi: Segundo argumento.
 - rdx: Tercer argumento.
 - rcx: Cuarto argumento.
 - r8: Quinto argumento.
 - r9: Sexto argumento.

### Almacenamiento de datos

Estos van en la sección `.data`, la sintaxis para cada dato es:

```nasm
section .data
    buffer db 100 ; Reservar espacio.
    hello db 'Hola Mundo!', 0 ; Definir cadena de caracteres.
```

Literalmente cada _byte_ va separado por una coma.

## Notas


### Descripción de los registros

 1. **RAX**: Registro de Acumulador. Se utiliza para operaciones aritméticas y de cálculo.

 2. **RBX**: Registro de Base. A menudo se utiliza como un puntero base para el acceso a datos en estructuras o matrices.

 3. **RCX**: Registro de Contador. Se utiliza en bucles y otras operaciones de control de flujo.

 4. **RDX**: Registro de Datos. Se utiliza para operaciones aritméticas y de datos, como división.

 5. **RSI**: Registro de Fuente Índice. A menudo se usa como un puntero para leer datos desde la memoria.

 6. **RDI**: Registro de Destino Índice. A menudo se usa como un puntero para escribir datos en la memoria.

 7. **RSP**: Registro de Pila. Mantiene la dirección actual de la cima de la pila, que se utiliza en la administración de la pila.

 8. **RBP**: Registro de Base de Pila. Se utiliza a menudo como un marco de pila base en funciones para acceder a las variables locales y parámetros.

 9. **R8 - R15**: Registros de Propósito General adicionales. Se utilizan para almacenar datos temporales y argumentos adicionales en llamadas a funciones.

 10. **RIP**: Registro de Instrucción de Puntero. Almacena la dirección de la siguiente instrucción que se ejecutará.

 11. **RFLAGS**: Registro de Banderas. Almacena diversas banderas de estado que indican condiciones como el desbordamiento, la igualdad, etc.

 12. **CS, DS, SS, ES, FS, GS**: Segmentos de registros. Controlan el acceso a segmentos de memoria en modo de segmentación, aunque en la arquitectura x86-64, se utilizan principalmente en el modo de compatibilidad y no son tan relevantes.

 13. **XMM0 - XMM15**: Registros de SIMD. Se utilizan para operaciones de instrucciones multimedia y vectoriales, como SSE y AVX.

 14. **ST0 - ST7**: Registros de punto flotante. Se utilizan para operaciones de punto flotante en la unidad de coma flotante.

 15. **MM0 - MM7**: Registros de MMX. Utilizados en instrucciones MMX para procesamiento de datos multimedia.

