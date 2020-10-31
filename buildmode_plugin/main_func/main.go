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
	plugin_NewLogger, err := plugin_log_os.Lookup("NewLogger")
	if err != nil {
		panic(err)
	}
	newLogger := plugin_NewLogger.(func(io.Writer) interface{})

	loggeri := newLogger(os.Stderr)
	fmt.Printf("%T\n", loggeri)

	logger := loggeri.(Log)
	logger.SetOut(os.Stderr)
	logger.Infof("this %s\n", "plugin log")
}
