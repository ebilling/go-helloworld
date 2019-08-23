package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

const (
	teststring = "A test string for the main tests"
)

func TestTheMainFunction(t *testing.T) {
	os.Args = []string{"test", teststring}
	orig, r, w := captureStdOut()
	main()
	out, err := readAllFromRedirectAndRestore(orig, r, w)
	assert.NoError(t, err)
	assert.Equal(t, teststring, string(out))
}

func captureStdOut() (*os.File, io.ReadCloser, io.WriteCloser) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	return rescueStdout, r, w
}

func readAllFromRedirectAndRestore(original *os.File, r io.ReadCloser, w io.WriteCloser) (out []byte, err error) {
	w.Close()
	out, err = ioutil.ReadAll(r)
	os.Stdout = original
	return out, err
}
