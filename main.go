package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var builder strings.Builder
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			builder.WriteString(strconv.Itoa(i)+",")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(builder.String())
}

type User struct {
	name    string
	number  int32
	number2 int
}

//completableFuture