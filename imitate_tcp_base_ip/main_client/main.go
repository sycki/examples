package main

import (
	"os"

	"github.com/sycki/examples/imitate_tcp_base_ip"
)

func main() {
	err := imitate_tcp_base_ip.DialTCP(os.Args[1], os.Args[2])
	if err != nil {
		println(err.Error())
	}
}
