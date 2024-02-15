package main

import (
	"fmt"
	"log"
	"os"
)

const (
	memsize = 30000
)

var (
	program []byte
	data    [memsize]byte
	p       uint16
	rp      int
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid argument count")
	}
	program = []byte(os.Args[1])
	read()
}

func incP() {
	if p >= memsize {
		fmt.Println(p)
		log.Fatal("invalid memory address")
	}
	p++
}

func decP() {
	if p <= 0 {
		fmt.Println(p)
		log.Fatal("invalid memory address")
	}
	p--
}

func inc() {
	data[p]++
}

func dec() {
	data[p]--
}

func print() {
	fmt.Printf("%c", data[p])
}

func read() {
	for rp < len(program) {
		b := program[rp]
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
			if data[p] == 0 {
				jmp()
			}
		case ']':
			if data[p] != 0 {
				revjmp()
			}
		default:
			log.Fatal("invalid character")
		}
		rp++
	}
}

func jmp() {
	nest := 0
	for rp < len(program)-1 && nest >= 0 {
		rp++
		b := program[rp]
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
	for rp > 0 && nest >= 0 {
		rp--
		b := program[rp]
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
