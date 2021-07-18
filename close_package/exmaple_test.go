package close_package

import (
	"fmt"
	"testing"
)

//函数片段
func add(base *int) func(int) int {

	fmt.Printf("%p\n", base)  //打印变量地址，可以看出来 内部函数时对外部传入参数的引用
	i := 10
	a := make([]int, 10)
	defer func() {
		//闭包内就是指针、随着加1
		fmt.Printf("%p , value : %v \n", base,*base)
		//30 跟随修改
		fmt.Println(i)
	}()
	//切片值复制，函数内会随着变化
	defer fmt.Printf("type : %T,value :%v \n",a,a)
	//一样的地址,没有*base没有变化还是10
	defer fmt.Printf("%p , value : %v \n", base,*base)
	*base++
	i = 30
	a=append(a, 1,2,3,4,5,6,7,8,9,10)

	f := func(i int) int {
		fmt.Printf("%p\n", base)
		*base += i
		return *base
	}

	return f
}

func TestAdd(t *testing.T) {
	base := 10
	t1 := add(&base)
	fmt.Println("-----------")
	base = 20
	fmt.Println(t1(1), t1(2))

	//t2 := add(100)
	//fmt.Println(t2(1), t2(2))
}

//github.com/TrueFurby/go0callvis
//github.com/cweill/gotests