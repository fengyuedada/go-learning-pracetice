package multiThreading

import (
	"fmt"
	"sync"
	"time"
)

var count2 int
var mutex sync.Mutex //定义一个锁的变量(传参一定要传地址，否则就会导致加锁无效)
func test1Mutex() {
	for i := 0; i < 10000; i++ {
		mutex.Lock() //对共享变量操作之前先加锁
		count2++
		defer mutex.Unlock() //对共享变量操作完毕在解锁，这样就保护了共享的资源
	}
}

func test2Mutex() {
	for i := 0; i < 10000; i++ {
		//mutex.Lock()
		count2++
		//defer mutex.Unlock()
	}
}

func TestMutex() {
	go test1Mutex()
	go test2Mutex()

	time.Sleep(time.Second)
	fmt.Println(count2) //输出结果是20000
}
