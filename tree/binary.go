package tree

import (
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	if t != nil {
		wg.Add(1)
		go Walk(t.Left, ch, wg)
		ch <- t.Value
		wg.Add(1)
		go Walk(t.Right, ch, wg)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(1)
	wg2.Add(1)
	go func() {
		Walk(t1, ch1, &wg)
		wg.Wait()
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2, &wg2)
		wg2.Wait()
		close(ch2)
	}()
	//两个类似Set比较数据结构
	ch1Map := make(map[int]int, 20)
	ch2Map := make(map[int]int, 20)
	for i := range ch1 {
		ch1Map[i] = ch1Map[i] + 1
	}
	for i := range ch2 {
		ch2Map[i] = ch2Map[i] + 1
	}
	if len(ch1Map) != len(ch2Map) {
		return false
	}
	for key, value := range ch1Map {
		if ch2Map[key] != value {
			return false
		}
		fmt.Println("key : ", key, " value : ", value)
	}
	return true
}
