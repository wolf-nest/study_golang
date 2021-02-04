package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		fmt.Print("hello world!", os.Args[1])
	}
}
