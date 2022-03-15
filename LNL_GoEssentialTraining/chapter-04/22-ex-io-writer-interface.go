package main

import (
	"fmt"
	"io"
	"os"
)

type Capitalizer struct {
	wtr io.Writer
}

func (c *Capitalizer) Write(p []byte) (n int, err error){
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, char := range p {
		if char >= 'a' && char <= 'z' {
			char -= diff
		}
		out[i] = char
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capitalizer{os.Stdout}
	_, err := fmt.Fprintln(c, "Hello there")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}