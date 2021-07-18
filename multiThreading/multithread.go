package multiThreading

import (
	"fmt"
	"time"
)

var count int

func test1() {
	for i := 0; i < 10000; i++ {
		count++
		// count += 1
		// count = count + 1
	}
}

func test2() {
	for i := 0; i < 10000; i++ {
		count++
		// count += 1
		// count = count + 1
	}
}

func TestMulti() {
	//runtime.GOMAXPROCS(2) //Go1.5版本之前，默认使用的是单核心执行，需要手动开启多核
	go test1()
	go test2()

	time.Sleep(time.Second)
	fmt.Println(count) //输出不稳定，一般会小于20000，预期值是20000
}

func TestAppend() {
	//runtime.GOMAXPROCS(2) //Go1.5版本之前，默认使用的是单核心执行，需要手动开启多核
	var arr []int

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				arr = append(arr, j)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(len(arr)) //输出不稳定，一般会小于1000，预期值是1000
}

func TestRange() {
	var s []int = []int{1, 2, 3}
	for _, value := range s {
		//value 的内存地址没有变过
		fmt.Printf("the value address is %p\n", &value)
	}
}

func TestRangeGo() {
	var arr []int
	// 定义一个数组并给它append 0~9的值
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	for _, v := range arr {
		go func() {
			fmt.Printf("v=%v head=%p\n", v, &v)
		}()
	}
	time.Sleep(time.Second)

	//输出结果
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//2
	//v= 9
	//v= 9
	//v= 9
	//v= 9
	//v= 8
	//v= 3
	//v= 9
	//v= 9
	//v= 9
	//v= 9
}

//加锁
