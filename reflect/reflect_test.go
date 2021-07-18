package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Pig struct {
	Weight int `json:"weight"`
	Mood string `json:"mood"`
}

func (p *Pig) Laugh(count int) (string,error) {
	for i := 0; i < count; i++ {
		fmt.Println(p.Mood," : hahaha")
	}
	return "done",nil
}


func TestReflect(t *testing.T) {
	//go对类型很严格
	pig := Pig{300, "happy"}
	of := reflect.ValueOf(pig)
	//不知道的情况查值
	name := of.FieldByName("Weight")
	t.Log(name)
	if byName, b := reflect.TypeOf(pig).FieldByName("Weight");b {
		t.Log(byName.Tag.Get("json"))
	}
	valueOf := reflect.ValueOf(&pig)
	call := valueOf.MethodByName("Laugh").Call([]reflect.Value{reflect.ValueOf(3)})
	t.Log(call)
	//知道对象的情况直接转
	p := of.Interface().(Pig)
	t.Log(p)
	t.Log(p.Laugh(3))
}

func TestFillWeight(t *testing.T) {
	settings := map[string]interface{}{"Weight": 300}
	e := Pig{}
	//e := 123
	err := fillBySettings(&e, settings)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(e)
}

func fillBySettings(st interface{}, settings map[string]interface{}) error{
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("error,not a ptr")
	}
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("error,not ptr to  a struct")
	}
	if settings == nil {
		return errors.New("settings is nil")
	}
	var(
		field reflect.StructField
		ok bool
	)
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st).Elem().Type().FieldByName(k)) ;!ok{
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st).Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

type int10086 int

func TestUnsafe(t *testing.T) {
	var i = int10086(1997)
	t.Logf("%T \n",i)//int
	i2 := *(*int)(unsafe.Pointer(&i))
	t.Logf("%T \n",unsafe.Pointer(&i))//unsafe.Pointer
	t.Logf("%T \n",int(i2))//float64
	t.Log(int(i2))//9.866e-321
}

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func(op int) {
		data := []int{}
		for i := 0; i < 100+op; i++ {
			data = append(data,i)
		}
		//shareBufPtr = unsafe.Pointer(&data)
		atomic.StorePointer(&shareBufPtr,unsafe.Pointer(&data))
	}
	readDatFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		//data := shareBufPtr
		fmt.Println(data,*(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn(0)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn(i)
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDatFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

