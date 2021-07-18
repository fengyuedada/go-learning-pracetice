package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const (
	row = 10000
	col = 10000
)

func fillMatrix(m *[row][col]int)  {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int)  {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("error create cpu.prof")
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal("error create cpu.prof")
	}
	defer pprof.StopCPUProfile()


	//主要逻辑
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	//mem输出
	f1, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("error create mem.prof")
	}
	//GC,获取最新的数据信息
	//runtime.GC()
	err = pprof.WriteHeapProfile(f1)
	if err != nil {
		log.Fatal("error write mem")
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("goroutine error")
	}

	gProf := pprof.Lookup("goroutine")
	if gProf == nil {
		log.Fatal("error gProf")
	}else {
		gProf.WriteTo(f2,0)
	}
	f2.Close()

}
