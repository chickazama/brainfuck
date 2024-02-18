package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	memsize = 32768
)

var (
	program        []byte
	memory         [memsize]byte
	memoryPtr      int
	instructionPtr int
	reader         *bufio.Reader
)

func init() {
	memoryPtr = 0
	instructionPtr = 0
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid argument count")
	}
	program = []byte(os.Args[1])
	execute()
}

func execute() {
	for instructionPtr < len(program) {
		b := program[instructionPtr]
		switch b {
		case '<':
			decP()
		case '>':
			incP()
		case '-':
			dec()
		case '+':
			inc()
		case '.':
			print()
		case '[':
			if memory[memoryPtr] == 0 {
				jmp()
			}
		case ']':
			if memory[memoryPtr] != 0 {
				revjmp()
			}
		case ',':
			b, err := reader.ReadByte()
			if err != nil {
				log.Fatal(err.Error())
			}
			memory[memoryPtr] = b
			reader.Reset(os.Stdin)
		}
		instructionPtr++
	}
}

func incP() {
	if memoryPtr >= memsize {
		fmt.Println(memoryPtr)
		log.Fatal("invalid memory address")
	}
	memoryPtr++
}

func decP() {
	if memoryPtr <= 0 {
		fmt.Println(memoryPtr)
		log.Fatal("invalid memory address")
	}
	memoryPtr--
}

func inc() {
	memory[memoryPtr]++
}

func dec() {
	memory[memoryPtr]--
}

func print() {
	fmt.Printf("%c", memory[memoryPtr])
}

func jmp() {
	fmt.Println(instructionPtr)
	nest := 0
	for instructionPtr < len(program)-1 && nest >= 0 {
		instructionPtr++
		b := program[instructionPtr]
		switch b {
		case '[':
			nest++
		case ']':
			nest--
		}
	}
	if nest >= 0 {
		log.Fatal("invalid program construct")
	}
}

func revjmp() {
	nest := 0
	for instructionPtr > 0 && nest >= 0 {
		instructionPtr--
		b := program[instructionPtr]
		switch b {
		case '[':
			nest--
		case ']':
			nest++
		}
	}
	if nest >= 0 {
		log.Fatal("invalid program construct")
	}
}
