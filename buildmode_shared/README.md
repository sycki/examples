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

## 通过动态库解耦（新版本可能已经不支持）
先将Go标准库编译为动态库，所有涉及的依赖库将被编译为.a
```
go install -buildmode=shared -linkshared std
ll -h $GOROOT/pkg/linux_amd64_dynlink/libstd.so
```

再将自己的包编译为动态库
```
go install -buildmode=shared -linkshared ./calculate
ll -h libgithub.com-sycki-examples-dynamic_lib-calculate.so
```
如果禁用了go mod模式，则生成的.so文件在`$GOPATH/pkg/linux_amd64_dynlink`目录中。

最后以动态链接方式编译main包，没有前两步同样可以执行以下操作，但会很慢，因为前面两步同样会被执行，且.so文件不会保留下来，相当于每次编译都会执行前两步。（1.14.6中貌似已经不支持，编译出来的main，没有以动态库方式引用calculate，把动态库删掉同样能运行，看样子只能用cgo调用动态库了）
```
go build -linkshared cmd/main.go
ll -h main
```

如果之后删除了动态库libstd.so，在编译自己的包时可能会报错。
```
readELFNote failed: open /home/sycki/program/go/pkg/linux_amd64_dynlink/libstd.so: no such file or directory
```

这时需要把动态库删掉，重新生成动态库。
```
rm -rf $GOROOT/pkg/linux_amd64_dynlink
```

