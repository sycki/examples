package main

import (
	"fmt"
	"io"
)

func init() {
	// Instance = &Logger{}
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
	// 返回对象必须是&Logger{}
	// 如果是Logger{}，则以下所有函数必须为 func (l Logger) 形式
	// 否则强转时报错main.Logger没有实现所有 func (l *Logger) 形式的函数
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
