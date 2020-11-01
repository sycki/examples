package main

type People struct{}

func (p People) GetName() string      { return "jack" }
func (p *People) SetName(name string) {}

type PeopleI interface {
	GetName() string
	SetName(string)
}

func main() {
	// 转换时会panic，因为People方法集不包含SetName
	var p1 interface{} = People{}
	p2 := p1.(PeopleI)
	println(p2.GetName())

	// 可以正常转换
	var p3 interface{} = &People{}
	p4 := p3.(PeopleI)
	println(p4.GetName())
}
