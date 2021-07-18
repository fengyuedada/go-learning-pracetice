package zifu

import "fmt"

func TestByte() {
	a := 'a'
	b := '你'
	var c byte = 'a'

	//自我推到默认是rune
	fmt.Printf("%d %T\n", a, a)
	fmt.Printf("%d %T\n", b, b)
	fmt.Printf("%d %T\n", c, c)
}
