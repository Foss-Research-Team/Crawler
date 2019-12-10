package main

import (
	"fmt"
)

func main() {

	var url []byte = []byte("swiss")

	url = url[2:]

	fmt.Printf("%s\n",url)
	

}
