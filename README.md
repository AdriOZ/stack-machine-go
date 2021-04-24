# stack-machine-c

Implementation of an interpreter for a simple stack machine language.

The language specification is very simple:

- `PUSH <number>`: pushes a number.
- `POP`: pops a number.
- `ADD`: pops two numbers, sums them and pushes the result.
- `IFEQ <line>`: examines the top number. If the value is 0, continues. Else, it jumps to the specified line number.
- `JUMP <line>`: jumps to the specified line number.
- `PRINTN`: prints the value at the top of the stack as a number.
- `PRINTC`: prints the value at the top of the stack as a number (it casts the number to an integer and prints the ASCII representation of it).
- `DUP`: pushes a duplicate of the value at the top of the stack.

To keep things simple, only one instruction will be interpreted per line.

Lines starting with `#` will be ignored (comments)
