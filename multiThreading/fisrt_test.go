package multiThreading

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOne(t *testing.T) {
	//t.Log("My first try")
	//a := 1
	//b := 1
	//t.Log(a,b)
	//for i := 0; i < 10; i++ {
	//	a = nextFibonacci(a, b)
	//	b = nextFibonacci(b,a)
	//	t.Log(a,b)
	//}
	//
	//i := 64
	//address := &i
	////指针不支持运算——访问连续空间
	//t.Log(address)
	//t.Logf("%T %T",*address,address)

	//i2 := [4]int{1, 2, 3}
	//i4 := [4]int{1, 2, 3}
	//for i3 := range i2 {
	//	t.Log(i3)
	//}
	////长度相等才能比较
	//t.Log(i2==i4)

	//&^按位清零，用来清除权限功能
	//i := 7
	//i=i&^Readable
	//i=i&^Writeable
	//
	//if i & Readable ==Readable{
	//	t.Log("is Readable")
	//}
	//if i & Writeable ==Writeable{
	//	t.Log("is Writeable")
	//}
	//if i & Executable ==Executable{
	//	t.Log("is Executable")
	//}
	//t.Log(on,off,appending)

	// 输出各数值范围
	//fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	//fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	//fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	//fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	//// 初始化一个32位整型值
	//var a int32 = 1047483647
	//// 输出变量的十六进制形式和十进制值
	//fmt.Printf("int32: 0x%x %d\n", a, a)
	//
	//// 将a变量数值转换为十六进制，发生数值截断
	//b := int16(a)
	//// 输出变量的十六进制形式和十进制值
	//fmt.Printf("int16: 0x%x %d\n", b, b)
	//
	//// 将常量保存为float32类型
	//var c float32 = math.Pi
	//// 转换为int类型，浮点发生精度丢失
	//fmt.Println(int(c))

	i := 1997
	address := &i
	fmt.Printf("num is : %d, it's location in memory: %p ,type is: %T \n", i, address, reflect.TypeOf(address))

}

func nextFibonacci(one int, two int) int {
	return one + two
}

const (
	on = 1 << iota
	off
	appending
)

//为了位运算
const (
	Readable = 1 << iota
	Writeable
	Executable
)
