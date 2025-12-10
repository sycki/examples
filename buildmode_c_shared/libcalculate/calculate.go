package main

import "C"

//export calculate
func calculate(methodC *C.char, aC C.int, bC C.int) C.int {
	method := C.GoString(methodC)
	a := int(aC)
	b := int(bC)

	var r int
	switch method {
	case "plus":
		r = a + b
	case "sub":
		r = a - b
	}

	return C.int(r)
}

func main() {}
