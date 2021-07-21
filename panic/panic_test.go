package panic

import (
	"fmt"
	"runtime"
	"testing"
)

func a() {
	fmt.Println("a")
	b()
}

func b() {
	fmt.Println("b")
	c()
}

type Student struct {
	Name int
}

func c() {
	defer RecoverFromPanic("fun c")
	fmt.Println("c")
	var a *Student
	fmt.Println(a.Name)
}

func TestPanic(t *testing.T) {
	a()
}

//func RecoverFromPanic(funcName string) {
//	if e := recover(); e != nil {
//		buf := make([]byte, 64<<10)
//		buf = buf[:runtime.Stack(buf, false)]
//
//		fmt.Printf("[%s] func_name: %v, stack: %s", funcName, e, string(buf))
//	}
//
//	return
//}

func RecoverFromPanic(funcName string) {
	if e := recover(); e != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]

		//logs.Errorf("[%s] func_name: %v, stack: %s", funcName, e, string(buf))

		panicError := fmt.Errorf("%v", e)
		ReportPanic(panicError.Error(), funcName, string(buf))
	}

	return
}
