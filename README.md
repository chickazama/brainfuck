# Brainfuck

A simple 'Brainfuck' interpreter, written in Go.

## Instructions

| Instruction | Description                                                                            |
| ----------- | -------------------------------------------------------------------------------------- |
| `<`         | Moves the memory pointer one cell to the left.                                         |
| `>`         | Moves the memory pointer one cell to the right.                                        |
| `+`         | Increments the current cell by one.                                                    |
| `-`         | Decrements the current cell by one.                                                    |
| `,`         | Reads one byte of data from stdin.                                                     |
| `.`         | Writes the value of the current cell to stdout as ASCII.                               |
| `[`         | Jumps forward to the instruction after the matching `]` if the current cell is 0.      |
| `]`         | Jumps backwards to the instruction after the matchin `[` if the current cell is not 0. |


## Useful Information

- Supply your program source as a command-line argument string.
- The '<', '>', '-', '+', '[', and ']' instructions are defined as expected.
- The ',' instruction implementation is to read in one byte (buffered by newline) from stdin.
- The '.' instruction implementation is to print the corresponding ASCII value (and values defined beyond ASCII up to code-point 255) to stdout.
- All other characters in the program source are ignored by the interpreter.
- Maximum Memory: 32kB (i.e. 32768 byte locations)
- Each cell in memory stores one byte (8-bit unsigned integer)
- Tape does not wrap (i.e. you cannot overflow memory locations on either side)
- Byte values do wrap (i.e. you can indefinitely increment/decrement byte values, and they will wrap round)

## Interesting Links

- [Brainfuck Reference](https://brainfuck.org/brainfuck.html)
- [Some Brainfuck Fluff (Daniel B. Cristofani)](https://brainfuck.org)
- [Urban Muller's Canonical Brainfuck Distribution](https://aminet.net/package.php?package=dev/lang/brainfuck-2.lha)


