package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	fmt.Println("修改前：", slice)

	slice2 := slice[:]
	slice2[0] = 100
	// 追加会生成新的切片
	// slice2 = append(slice2, 3)
	fmt.Println("修改后：", slice)
}
