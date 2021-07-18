package tree

import (
	"fmt"
	"golang.org/x/tour/tree"
	"strconv"
	"testing"
	"time"
)



func TestOneMain(t *testing.T) {
	same := Same(tree.New(1), tree.New(1))
	t.Log(same)
	t.Log(Same(tree.New(1), tree.New(2)))
}

type People struct {
	name string
	age  int
}

func TestMap(t *testing.T) {
	m1 := map[string]People{"1": {name: "张三", age: 30}}
	people, ok := m1["1"]
	if ok {
		t.Log(people)
	} else {
		t.Log("not existing", people)
	}
}

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * 2 }
	m[3] = func(op int) int { return op * 3 }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mSet := map[int]bool{}
	//添加元素
	mSet[1] = true
	n := 1
	//查找
	if mSet[n] {
		t.Logf("%d is exist", n)
	} else {
		t.Logf("%d is not exist", n)
	}
	//查找元素个数
	t.Log(len(mSet))
	//删除
	delete(mSet, 1)
}

func timeSpent(inner func(op int) string) func(op int) string {
	return func(n int) string {
		start := time.Now()
		//前后封一层时间然后返回
		ret := inner(n)
		fmt.Println("time Spent :", time.Since(start).Seconds())
		return ret
	}
}

func slowProductionPig(op int) string {
	time.Sleep(time.Duration(op) * time.Second)
	return strconv.Itoa(op) + "只猪"
}

func TestName(t *testing.T) {
	//函数式编程，无感知
	productPig := timeSpent(slowProductionPig)
	t.Log(productPig(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestNarparam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 2, 3, 4))
}

func (p *People) String1() string {
	fmt.Printf("%X \n", p)
	*p = People{"张三", 29}
	fmt.Println("string1 :name", &p.name)
	p.name = "王五"
	return "str is " + p.name + strconv.Itoa(p.age)
}

func (p People) String2() string {
	fmt.Printf("%X \n", &p)
	fmt.Println("string2 name:", &p.name)
	p.name = "李四"
	p = People{"", 0}
	return "str is " + p.name + strconv.Itoa(p.age)
}

func TestPeopleFun(t *testing.T) {
	people := &People{"张三", 30}
	t.Logf("%X", &people.name)
	t.Log(people.String1())
	t.Log(people.String2())
	t.Log(people)
}

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	//改用匿名变量
	//p *Pet

	//类似继承，可以获得相应方法，实际更像组合
	Pet
}

//无法重载
func (p *Dog) Speak() {
	fmt.Print("Wang")
}

func TestDog(t *testing.T) {
	//不支持LSP，dog类型当Pet用
	dog := new(Dog)
	dog.SpeakTo("Chao")
}

//再次提醒有没有等于的区别
type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Printlm(\"Hello World\")"
}

func (g *JavaProgrammer) WriteHelloWorld() Code {
	return "sys.out.println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	//接口只能接受指针，下面这两个都是指针，也可以写成&GoProgrammer{}
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
	t.Logf("%T", Code("2"))
}
