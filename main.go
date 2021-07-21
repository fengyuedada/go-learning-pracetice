package main

import (
	"fmt"
	"os"
)

func main() {
	//var builder strings.Builder
	//var wg sync.WaitGroup
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)
	//	go func() {
	//		builder.WriteString(strconv.Itoa(i)+",")
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
	//fmt.Println(builder.String())
	fmt.Println(os.Getenv("HOSTNAME"))
}

type User struct {
	name    string
	number  int32
	number2 int
}

//completableFuture
