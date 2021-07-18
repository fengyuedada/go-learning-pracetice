package slice

import (
	"flag"
	"fmt"
	"strings"
)

//定义一个类型，用于增加该类型方法
type sliceValue []string

//new一个存放命令行参数值的slice
func newSliceValue(p *[]string) *sliceValue {
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = strings.Split(val, ",")
	return nil
}

//flag为slice的默认值default is me,和return返回值没有关系
func (s *sliceValue) String() string {
	*s = strings.Split("default is me", ",")
	//run only one
	return "defValue useless"
}

func main() {
	var languages []string
	flag.Var(newSliceValue(&languages), "slice", "I like programming `languages`")
	var ageInt int64
	age := flag.Int64("age", 0, "people age")
	flag.Int64Var(&ageInt, "age1", 0, "people age")
	flag.Parse()
	fmt.Println(flag.Lookup("slice").DefValue)

	//打印结果slice、age接收到的值
	fmt.Println(languages)
	fmt.Println(*age)
	fmt.Println(ageInt)
}
