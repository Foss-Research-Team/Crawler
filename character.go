package main

import (
	"fmt"
)

func main() {

	b := byte(0b00010011)
    	fmt.Printf("%08b %02x\n", b, b)
    	x := byte(0x13)
    	fmt.Printf("%08b %02x\n", x, x)

	

}
