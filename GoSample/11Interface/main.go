/*
	Interface co kha nang compose
*/
package main

import (
	"bytes"
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
cach xai con tro
func (cw *ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
*/

func main() {
	fmt.Println("Interface ")

	//var w Writer = &ConsoleWriter{}

	var w Writer = ConsoleWriter{}
	w.Write([]byte("hello world"))

	// cach 2 tao interface
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 1; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriteCloser = NewBufferWriterCloser()
	wc.Write([]byte("hello world - Hello World"))
	wc.Close()
}

// cach 2 tao interface
type Incrementer interface {
	Increment() int
}

// cach 2 tao interface
type IntCounter int

// cach 2 tao interface
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/////////////////////
// Interface co kha nang compose

type Write interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}
type WriteCloser interface {
	Writer
	Closer
}
type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 0)
	for bwc.buffer.Len() > 0 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Printf(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(0)
		_, err := fmt.Printf(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
