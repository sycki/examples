package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	listen()
}

func listen() error {
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer terminal.Restore(fd, state)

	b := make([]byte, 128)
	for {
		l, err := os.Stdin.Read(b[:])
		if err != nil {
			if err == io.EOF {
				return nil
			}
		}

		fmt.Printf("len [%d] byte %d rune %c\r\n", l, b[:l], b[:l])

		if l == 1 && b[0] == 4 {
			return nil
		}
	}
}
