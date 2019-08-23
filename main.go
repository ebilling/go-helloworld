package main

import (
	"fmt"
	"os"

	"github.com/ebilling/go-helloworld/hello"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"STRING YOU WANT TO PRINT\"\n", os.Args[0])
		os.Exit(1)
	}
	h := hello.NewHello(os.Args[1])
	h.Print()
}
