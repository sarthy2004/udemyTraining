package main

import (
	"fmt"
)

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %c\n", 42, 42)
	fmt.Printf("%d - %c - %x - %X - %#X \n", 42, 42,42, 42,42)
	
	for i:=0; i < 200; i++{
	fmt.Printf("%d - %c - %x - %X - %#X \n", i, i, i,i,i)
	}
}
