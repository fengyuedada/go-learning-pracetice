package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok := p.(int);ok {
	//	fmt.Println("int :",i)
	//	return
	//}
	//
	//if s,ok := p.(string);ok {
	//	fmt.Println("string :",s)
	//	return
	//}
	//fmt.Println("Unknown Type",p)
	switch v := p.(type) {
	case int:
		fmt.Println("int :", v)
	case string:
		fmt.Println("string :", v)
	case func():
		fmt.Println("fun :", &v)
	default:
		fmt.Println("Unknown Type", v)
	}
}

func TestInterface(t *testing.T) {
	DoSomething(func(p string) {})
}
