package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	memsize      = 32768
	expectedArgc = 2
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
	if len(os.Args) != expectedArgc {
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
			decrementMemoryPtr()
		case '>':
			incrementMemoryPtr()
		case '-':
			decrementByte()
		case '+':
			incrementByte()
		case '.':
			print()
		case '[':
			if memory[memoryPtr] == 0 {
				jumpForward()
			}
		case ']':
			if memory[memoryPtr] != 0 {
				jumpBackward()
			}
		case ',':
			readByte()
		}
		instructionPtr++
	}
}

func decrementMemoryPtr() {
	if memoryPtr <= 0 {
		fmt.Println(memoryPtr)
		log.Fatal("invalid memory address")
	}
	memoryPtr--
}

func incrementMemoryPtr() {
	if memoryPtr >= memsize {
		fmt.Println(memoryPtr)
		log.Fatal("invalid memory address")
	}
	memoryPtr++
}

func decrementByte() {
	memory[memoryPtr]--
}

func incrementByte() {
	memory[memoryPtr]++
}

func print() {
	fmt.Printf("%c", memory[memoryPtr])
}

func jumpForward() {
	fmt.Println(instructionPtr)
	nestLevel := 0
	for instructionPtr < len(program)-1 && nestLevel >= 0 {
		instructionPtr++
		b := program[instructionPtr]
		switch b {
		case '[':
			nestLevel++
		case ']':
			nestLevel--
		}
	}
	if nestLevel >= 0 {
		log.Fatal("invalid program construct")
	}
}

func jumpBackward() {
	nestLevel := 0
	for instructionPtr > 0 && nestLevel >= 0 {
		instructionPtr--
		b := program[instructionPtr]
		switch b {
		case '[':
			nestLevel--
		case ']':
			nestLevel++
		}
	}
	if nestLevel >= 0 {
		log.Fatal("invalid program construct")
	}
}

func readByte() {
	b, err := reader.ReadByte()
	if err != nil {
		log.Fatal(err.Error())
	}
	memory[memoryPtr] = b
	reader.Reset(os.Stdin)
}
