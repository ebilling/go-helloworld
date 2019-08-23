package hello

import (
	"io"
	"os"
)

// Hello is the dumbest object ever, it just holds a string and you can print it to stdout
type Hello struct {
	text string
	out  io.Writer
}

// Intentional Linting issue - first line should be the second line here
// NewHello creates a Hello with the given string
func NewHello(s string) *Hello {
	return &Hello{
		text: s,
		out:  os.Stdout,
	}
}

func (h *Hello) setDestination(f io.Writer) {
	h.out = f
}

// Print prints the string to stdout
func (h *Hello) Print() {
	h.out.Write([]byte(h.text))
}
