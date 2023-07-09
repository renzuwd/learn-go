package main

import "fmt"

type Per struct{}

// func (p Per) Say() {
// 	fmt.Println("say")
// }

func (p *Per) Say() {
	fmt.Println("say")
}

// func (p *Per) Say() {}

func main() {
	// 引用类型初始化 会带上非引用类型方法
	var p *Per
	var p2 Per
	p.Say()
	(p2).Say()
}
