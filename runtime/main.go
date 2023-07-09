package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 查看 cpu 核心数
	fmt.Println(runtime.NumCPU())

	//设置golang运行的核心数
	runtime.GOMAXPROCS(runtime.NumCPU())
}
