package main

import "fmt"

type People struct{ name string }

func (p People) GetName() string      { return "jack" }
func (p *People) SetName(name string) {}

type Jack struct{ name string }

func (p Jack) GetName() string      { return "jack" }
func (p *Jack) SetName(name string) {}

type PeopleI interface {
	GetName() string
	SetName(string)
}

func main() {
	// 类型断言，i必须是interface{}，否则编译不通过，成功
	var i interface{} = &People{}
	pi := i.(PeopleI)
	println(pi.GetName())

	// 安全的类型断言，i必须是interface{}，否则编译不通过，成功
	if pi, ok := i.(PeopleI); !ok {
		fmt.Printf("failed i.(PeopleI): %T\n", i)
	} else {
		println(pi.GetName())
	}

	// 安全的类型断言，i必须是interface{}，且i的原型必须与目标类型相同才返回ok，失败
	if pi, ok := i.(*Jack); !ok {
		fmt.Printf("failed i.(*Jack): %T\n", i)
	} else {
		println(pi.GetName())
	}

	// 强制转换，p可以是任意类型，但必须与目标类型兼容才能转换，成功
	var p *People = &People{}
	pi = (*Jack)(p)
	println(pi.GetName())

}
