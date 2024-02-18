# Brainfuck

A simple 'Brainfuck' interpreter, written in Go.

## About

- Supply your program source as a command-line argument string.
- The '<', '>', '-', '+', '[', and ']' instructions are defined as expected.
- The ',' instruction implementation is to read in one byte (buffered by newline) from stdin.
- The '.' instruction implementation is to print the corresponding ASCII value (and values defined beyond ASCII up to code-point 255) to stdout.
- All other characters in the program source are ignored by the interpreter.
- Maximum Memory: 32kB (i.e. 32767 byte locations)
- Tape does not wrap (i.e. you cannot overflow memory locations on either side)
- Byte values do wrap (i.e. you can indefinitely increment/decrement byte values, and they will wrap round)

## Interesting Links

[Brainfuck Reference](https://brainfuck.org/brainfuck.html)
[Some Brainfuck Fluff (Daniel B. Cristofani)](https://brainfuck.org)


