package hello

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	teststring := "This is my test string"
	h := NewHello(teststring)
	assert.Equal(t, teststring, h.text)

	out := bytes.NewBuffer(nil)
	h.setDestination(out)
	h.Print()
	assert.Equal(t, teststring, out.String())
}

func TestSkipped(t *testing.T) {
	t.Skip("This is just a skipped test")
}

func TestFailedTest(t *testing.T) {
	assert.NotEqual(t, "true", os.Getenv("DO_FAIL"), "Failing because environment variable was \"true\"")
}
