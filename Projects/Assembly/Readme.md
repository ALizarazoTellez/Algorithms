
# Assembly Practice

Práctica en lenguaje ensamblador. Aquí añadiré los conceptos que vaya aprendiendo.

> **ADVERTENCIA**: Esto solo aplica para **Linux AMD64 (NASM)**.

## Resumen

Si estás haciendo un programa principal (es decir, el equivalente a `main` en otros lenguajes), recuerda llamar siempre a la función `exit`, porque de lo contrario el S.O. le enviará la señal 11 al programa (_segmentation fault_).

### Funciones de ensamblador

Cada llamada al sistema tiene un número asignado, este número se debe enviar al registro **rax**, después de le envían los argumentos en el correspondiente registro y se ejecuta **syscall**.

#### Argumentos de funciones

 - rdi: Primer argumento.
 - rsi: Segundo argumento.
 - rdx: Tercer argumento.
 - rcx: Cuarto argumento.
 - r8: Quinto argumento.
 - r9: Sexto argumento.

### Funciones útiles

#### Función `read`

 - Número de llamada al sistema: 0 (para read).
 - Primer argumento: Descriptor de archivo (ejemplo: stdin para entrada estándar).
 - Segundo argumento: Dirección de memoria donde se almacenarán los datos leídos.
 - Tercer argumento: Cantidad de bytes a leer.

#### Función `write`

 - Registro rax contiene el número de llamada al sistema para write, que es 1.
 - Registro rdi contiene el descriptor de archivo (file descriptor).
 - Registro rsi contiene la dirección de la memoria donde se encuentra el mensaje que se va a escribir.
 - Registro rdx contiene la longitud del mensaje.

#### Función `exit`

 - Número de llamada al sistema: 60 (para exit).
 - Primer argumento: Código de salida (un valor entero que indica el estado de salida).

Por alguna razón las personas prefieren usar `xor rdi, rdi` (en **XOR**, cuando dos booleanos son el mismo, da falso, que en ensamblador es cero), en lugar de `mov rdi, 0`.

### Almacenamiento de datos

Estos van en la sección `.data`, la sintaxis para cada dato es:

```nasm
    section .data

buffer:
    db 100 ; Reservar espacio.

hello:
    db 'Hola Mundo!', 0 ; Definir cadena de caracteres.
    hello_size equ $ - hello ; Tamaño de la cadena «hello».
```

Literalmente cada _byte_ va separado por una coma.

Por alguna razón que ignoro, no se puede usar la sintaxis de dos puntos para definir etiquetas con `equ`...

Una breve diferencia entre `db` y `equ`:


1. **`db` (Define Byte):**
   - Propósito: `db` se utiliza para definir y reservar espacio para datos en memoria. Es comúnmente utilizado para declarar constantes, como cadenas de caracteres, bytes individuales o matrices de bytes.
   - Ejemplo:
     ```nasm
     my_byte db 42      ; Define una constante de byte con valor 42
     my_string db 'Hello, world!',0  ; Define una cadena de caracteres terminada en null
     ```

2. **`equ` (EQUate):**
   - Propósito: `equ` se utiliza para definir constantes simbólicas. No reserva espacio en memoria, pero asigna un valor constante a un símbolo que puede ser utilizado en otras partes del código para mayor claridad y mantenimiento. Es una directiva utilizada en tiempo de ensamblaje para reemplazar un valor simbólico por su valor constante.
   - Ejemplo:
     ```nasm
     MY_CONSTANT equ 42     ; Define una constante simbólica llamada MY_CONSTANT con valor 42

     ; Uso de la constante simbólica
     mov rax, MY_CONSTANT  ; Equivale a mov rax, 42
     ```

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

