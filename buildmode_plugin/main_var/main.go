package main

import (
	"fmt"
	"io"
	"os"
	"plugin"
)

type Log interface {
	SetOut(out io.Writer)
	Infof(format string, a ...interface{})
	Infoln(msg interface{})
}

func main() {
	plugin_log_os, err := plugin.Open("./plugin_log.so")
	if err != nil {
		panic(err)
	}
	plugin_Instance, err := plugin_log_os.Lookup("Instance")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", plugin_Instance)
	logger := plugin_Instance.(Log)

	logger.SetOut(os.Stderr)
	logger.Infof("this %s\n", "plugin log")
}
