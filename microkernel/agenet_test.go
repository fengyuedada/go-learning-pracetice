package microkernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}


func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collector", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name)
	for {
		select {
		//上下文终止
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			fmt.Println("发送",c.name)
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector", c.name)
	select {
	//终止信号
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second * 1):
		return errors.New("failed to stop for timeout")
	}
}

func (c *DemoCollector) Destroy() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("c1", "1")
	c2 := NewCollect("c2", "2")
	agt.RegisterCollector("c1", c1)
	agt.RegisterCollector("c2", c2)
	agt.RegisterCollector("c3", c2)
	if err := agt.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	t.Log(agt.Start())
	time.Sleep(time.Second * 1)
	agt.Stop()
	agt.Destroy()
}
