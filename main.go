package main

import (
	"fmt"
	"log"
	"os"
)

const (
	memsize = 2048
)

var (
	program []byte
	data    [memsize]byte
	p       uint16
	rp      int
	fp      *os.File
)

func init() {
	var err error
	fp, err = os.OpenFile("expanded.txt", os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid argument count")
	}
	program = []byte(os.Args[1])
	read()
	fp.Close()
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
			fmt.Fprintf(fp, "%c", b)
		case '>':
			incP()
			fmt.Fprintf(fp, "%c", b)
		case '-':
			dec()
			fmt.Fprintf(fp, "%c", b)
		case '+':
			inc()
			fmt.Fprintf(fp, "%c", b)
		case '.':
			print()
			fmt.Fprintf(fp, "%c", b)
		case '[':
			if data[p] == 0 {
				jmp()
				// fp.WriteString("jmp\n")
			}
		case ']':
			if data[p] != 0 {
				revjmp()
				// fp.WriteString("revjmp\n")
			}
		default:
			log.Fatal("invalid character")
		}
		rp++
	}
}

func jmp() {
	fmt.Println(rp)
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
