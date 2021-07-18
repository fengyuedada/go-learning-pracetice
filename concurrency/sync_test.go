package concurrency

import (
	"context"
	"fmt"
	cm "github.com/easierway/concurrent_map"
	"sync"
	"testing"
	"time"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

func TestCounter(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	//time.Sleep(1*time.Second)
	wg.Wait()
	t.Logf("counter = %d", counter)
}

func TestSync(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Printf("%p\n", &ctx)
	fmt.Printf("%p\n", ctx)
	for i := 0; i < 10; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 1000000)
			}
			fmt.Println("caceled :", i)
		}(i, ctx)
	}
	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 1)
}

func cancel(cancelChan chan struct{}) {
	fmt.Println("主动关闭")
	close(cancelChan)
}

func isCanceled(ctx context.Context) bool {
	fmt.Printf("canceled  %p\n", &ctx)
	fmt.Printf("canceled  %p\n", ctx)
	select {
	case <-ctx.Done():
		return true
	default:
		return false
		//case <-time.After(time.Microsecond*100):
		//	fmt.Println("被动关闭")
		//	return true
	}
}
