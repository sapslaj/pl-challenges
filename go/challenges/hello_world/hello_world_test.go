package helloworld_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	helloworld "github.com/sapslaj/pl-challenges/go/challenges/hello_world"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, "hello world", helloworld.HelloWorld())
}

func TestPuts(t *testing.T) {
	// messing with os.Stdout is "bad" but it's a good example of how to capture
	// stdout when you have no control over the internal logic to pass an
	// io.Writer.
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	r, w, err := os.Pipe()
	assert.NoError(t, err)
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		assert.NoError(t, err)
		outC <- buf.String()
	}()

	helloworld.Puts()

	w.Close()
	os.Stdout = old
	assert.Equal(t, "hello world\n", <-outC)
}
