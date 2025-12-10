## C动态库使用
用go编写一个动态库，供c调用。

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

## 编译动态库
```
go build -o _output/libcalculate.so -buildmode=c-shared ./libcalculate
```

检查生成的动态库和头文件，动态库要以lib开头，否则gcc搜索不到。
```
ll _output/
libcalculate.h
libcalculate.so
```

检查动态库中是否有目标函数。
```
nm -D _output/libcalculate.so | grep ' T ' | grep -v _cgo
000000000008ca08 T crosscall1
000000000008bd60 T crosscall2
000000000008bff0 T fatalf
000000000008bea0 T calculate
```

## 在c中调用动态库
-L指定搜索路径，-l指定动态库名称，要去掉动态库文件的lib前缀和so后缀。
```
cp _output/libcalculate.h cmd/
gcc -o ./_output/calc ./cmd/main.c -L _output -l calculate
```

## 执行c程序
```
LD_LIBRARY_PATH=_output ./_output/calc
```
