package main

import (
	"fmt"
)

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b", 42, 42)
	fmt.Printf("%d - %c\n", 42, 42)
	fmt.Printf("%d - %c - %x - %X - %#X \n", 42, 42,42, 42,42)
}
