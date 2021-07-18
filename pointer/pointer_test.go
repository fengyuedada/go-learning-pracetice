package pointer

import (
	"fmt"
	"testing"
)

type T struct{
	Value 		int
}

func TestPointer(t *testing.T) {
	myT := T{Value:666}

	changeAd(&myT)
	t.Log(myT.Value)

	change(&myT)
	t.Log(myT.Value)
}
func changeAd(t *T){
	fmt.Printf("copy poniter(%p) has another adress %p \n",t,&t)
	t = &T{Value:999}
}
func change(t *T){
	*t = T{Value:999}
}
