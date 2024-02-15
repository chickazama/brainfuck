package main

import "fmt"

const (
	memsize = 2048
)

var (
	data [memsize]byte
	p    uintptr
)

func main() {
	for i := 0; i < memsize; i++ {
		data[i] = byte(i % 256)
	}
	fmt.Println(data[p])
}
