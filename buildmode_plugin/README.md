# 动态库使用
Golang在1.5版本中支持了生成动态库，1.8版本支持了插件功能。在C语言中，各模块之间可以通过二进制形式解耦，避免某一模块更新，其它模块必须重新导入新代码并重新编译。在Go语言1.8版本中也可以做到这一点，且优雅高效，他就是Go插件（plugin），缺点是Go编译出来的插件不能被C和其它语言直接使用，只兼容Go语言（这一点需要确认，官方没有说明）。

## 编译选项
```
go help buildmode
-buildmode=default      将main包编译为可执行文件，非main包编译为go兼容的.a静态库中
-buildmode=plugin       给定一个main包，编译出只有go兼容的.so动态库，使用"plugin"包加载和调用动态库中的函数
-buildmode=archive      给定一个非main包，将所有main依赖的包编译到一个go兼容的.a静态库中
-buildmode=shared       给定一个非main包，将所有main依赖的包编译到一个go兼容的.so动态库中
-buildmode=c-archive    给定一个main包，将main包以及所有依赖包编译为C兼容的.a静态库，只有那些//export的函数能够被调用
-buildmode=c-shared     给定一个main包，将main包以及所有依赖包编译为C兼容的.so动态库，只有那些//export的函数能够被调用
-buildmode=exe          给定一个main包，编译出exe可执行文件
-buildmode=pie          给定一个main包，编译出pie可执行文件
```

## 通过插件解耦
将plugin_log编译为插件动态库
```
go build -buildmode=plugin ./plugin_log
```

调用动态库中的函数
```
go run ./main_func/main.go
```

使用动态库中的结构体
```
go run ./main_var/main.go
```
