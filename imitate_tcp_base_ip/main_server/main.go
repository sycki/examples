package main

import (
	"os"

	"github.com/sycki/examples/imitate_tcp_base_ip"
)

func main() {
	err := imitate_tcp_base_ip.ListenTCP("127.0.0.1", os.Args[1])
	if err != nil {
		println(err.Error())
	}
}
