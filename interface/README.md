# 接口实现和类型转换
Go规范和原则：
- 结构体中的每个函数，都有一个接收者（receiver），在以下两个函数中，GetName的接收者是People，而SetName的接收者是`*People`。
- People的函数集中，包含一个函数，GetName，`*People`的函数集中包含两个函数，GetName和SetName。
- 调用原则，当`*People`可寻址且`*People`的函数集中包含SetName时，则`(*People).SetName`可以简写为`People.SetName`。

```
type People struct{}
func (p People) GetName() string      { return "jack" }
func (p *People) SetName(name string) {}

type PeopleI interface {
	GetName() string
	SetName(string)
}
```

## 接口实现
这时以下代码将不能通过编译，因为People的函数集中不包含SetName函数，故People没有实现PeopleI。
```
var p PeopleI = People{}
```

而以下代码可以正常运行，因为`*People`的函数集包含SetName和GetName，故`*People`实现了PeopleI。
```
var p PeopleI = &People{}
```

## 类型转换
Go中的类型转换大致可以分为两种，一种是强制转换，需要转换的对象与目标类型兼容即可转换。另一种是类型断言，用于判断指定的interface{}对象是否为某种类型，也就是说需要转换的对象必需是interface{}，且该对象的原型必须与目标相同才能完成转换。

现在定义一个与People兼容的类型Jack
```
type Jack struct{ name string }
func (p Jack) GetName() string      { return "jack" }
func (p *Jack) SetName(name string) {}
```

### 强制转换
因为`*People`与`*Jack`类型兼容，所以可以完成转换。注意因为`*`操作符优先级低于转换操作，所以要括起来`(*Jack)`。
```
pi = (*Jack)(&People{})
```

同样，因为byte与int兼容，所以可以完成转换。
```
var b byte = 100
n := int(b)
```

### 类型断言
因为i的原型是`*People`，所以断言成功。
```
var i interface{} = &People{}
pi, ok := i.(People)
```

因为i的原型是`*People`，且`*People`实现了PeopleI，所以断言成功。
```
var i interface{} = &People{}
pi, ok := i.(PeopleI)
```

因为i的原型是不是`*Jack`，所以断言失败。
```
var i interface{} = &People{}
pi, ok := i.(*Jack)
```
