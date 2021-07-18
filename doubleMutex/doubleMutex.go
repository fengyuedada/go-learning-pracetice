package doubleMutex

import (
	"fmt"
	"sync"
	"time"
)

var token chan int64
var mutex sync.Mutex //定义一个锁的变量(传参一定要传地址，否则就会导致加锁无效)

func getToken(name int) int64 {
	fmt.Println(mutex)
	if nil == token {
		mutex.Lock()
		if token == nil {
			time.Sleep(time.Second) //模拟耗时操作
			fmt.Println(name, "create token success")
			token = make(chan int64, 10)
			token <- 1
			token <- 2
			token <- 3
			token <- 4
			token <- 5
			fmt.Println(len(token))
		}
		mutex.Unlock()
	}
	i := <-token
	fmt.Println(name, "return token ", i, " least ", len(token))
	return i
}

func Test() {
	//runtime.GOMAXPROCS(2) //Go1.5版本之前，默认使用的是单核心执行，需要手动开启多核
	go getToken(1)
	go getToken(2)
	go getToken(3)
	go getToken(4)
	go getToken(5)
	time.Sleep(time.Second * 50)
}

func Setup2() <-chan bool {
	time.Sleep(time.Second * 3)
	c := make(chan bool)
	c <- true
	return c
}

func main() {
	if <-Setup2() {
		fmt.Println("setup succeed")
	}
}
