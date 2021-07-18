package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

var Method func() bool
var count = 2

//数据的生产者
func Producer(header string, channel chan<- string) {

	//不断的生产
	for {
		//将随机数和字符串格式化为字符串发给通道
		channel <- fmt.Sprintf("%s:%v", header, rand.Int31())

		time.Sleep(time.Second)

	}
}

//消费者
func Customer(channel <-chan string) {
	//获取数据
	for s := range channel {
		fmt.Println(s)
	}
}
