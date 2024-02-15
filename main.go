package main

import (
	"fmt"
	"log"
)

const (
	memsize = 2048
)

var (
	data [memsize]byte
	p    uint16
)

func main() {
	for i := 0; i < memsize; i++ {
		data[i] = byte(i % 256)
	}
	for i := 0; i < 65; i++ {
		inc()
	}
	print()
}

func incP() {
	if p >= memsize {
		log.Fatal("invalid memory address")
	}
	p++
}

func decP() {
	if p <= 0 {
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
