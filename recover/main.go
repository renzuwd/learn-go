package main

import (
	"fmt"
	"time"
)

func work() {
	panic("work is panic")
}

func process() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	work()
}

func main() {
	for {
		process()
		time.Sleep(time.Second)
	}
}
