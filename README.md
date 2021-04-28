# stack-machine-go

Implementation of an interpreter for a simple stack machine language.

The language specification is very simple:

**Stack Manipulation**

- `PUSH <number>`: pushes a number.
- `POP`: pops a number.
- `DUP`: pushes a duplicate of the value at the top of the stack.
- `SWAP`: pops two numbers and psuhes them back in reverse order.

**Operations**

- `ADD`: pops two numbers, sums them and pushes the result.
- `SUB`: pops two numbers, substracts them and pushes the result.
- `MUL`: pops two numbers, multiplies them and pushes the result.
- `DIV`: pops two numbers, divides them and pushes the result.

**Control flow**
- `IFEQ <line>`: examines the top number. If the value is 0, continues. Else, it jumps to the specified line number.
- `JUMP <line>`: jumps to the specified line number.

**Printing options**

- `PRINTN`: prints the value at the top of the stack as a number.
- `PRINTI`: prints the value at the top of the stack as an integer.
- `PRINTC`: prints the value at the top of the stack as a character (it casts the number to an integer and prints the ASCII representation of it).

To keep things simple, only one instruction will be interpreted per line.

Lines starting with `#` will be ignored (comments)
