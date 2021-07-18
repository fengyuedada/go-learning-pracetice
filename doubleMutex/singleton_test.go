package doubleMutex

import (
	meTree"awesomeProject/tree"
	"fmt"
	"golang.org/x/tour/tree"
	"sync"
	"testing"
)

type Singleton struct {
}
var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("first do")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestSyncOnce(t *testing.T) {
	var  wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			t.Logf("%p\n",obj)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new cache")
			return 100
		},
	}
	i := pool.Get().(int)
	fmt.Println(i)
	pool.Put(3)
	pool.Put(3)
	//runtime.GC() 	会清理pool
	i1,_ := pool.Get().(int)
	fmt.Println(i1)
	i2,_ := pool.Get().(int)
	fmt.Println(i2)
}

func BenchmarkNameTest(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
		meTree.Same(tree.New(1),tree.New(2))
		meTree.Same(tree.New(3),tree.New(3))
	}
	b.StopTimer()
	//性能无关测试代码
}