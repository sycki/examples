package main

import (
	"fmt"
	"io"
)

func init() {
}

// Lookup函数返回值类型是**main.Logger，导致无法强转为Log，因为指针的指针不会满足任何接口，除了interface{}
// var Instance = &Logger{}

//Lookup函数返回值类型是*main.Logger
var Instance Logger

type Logger struct {
	out io.Writer
}

// 返回值类型必须是所有人公共的类型，不能是本包独有的类型Logger，否则调用方无法定义和强转该函数
func NewLogger(out io.Writer) interface{} {
	// 如果返回对象是&Logger{}，则函数的接收者可以是*Logger也可以是Logger
	// 如果是Logger{}，则以下函数的接收者必须是Logger
	// 否则转换为Log时报错，因为Logger的函数集不包含以下函数
	return &Logger{
		out: out,
	}
}

func (l *Logger) SetOut(out io.Writer) {
	l.out = out
}

func (l *Logger) Infof(format string, a ...interface{}) {
	fmt.Fprintf(l.out, format, a...)
}

func (l *Logger) Infoln(msg interface{}) {
	fmt.Fprintln(l.out, msg)
}
