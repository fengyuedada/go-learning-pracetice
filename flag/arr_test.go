package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArr(t *testing.T) {
	var x = []int{1, 2, 3}
	var y = x
	fmt.Println(x, y)
	y[0] = 999
	fmt.Println(x, y)

	//数组的切片
	z := x[2:]
	z[0] = 4
	fmt.Println(x, z, cap(z))
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2), reflect.TypeOf(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2, year)
}
