package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i:=0; i<5; i++ {
		runtime.Gosched() //让出CPU的时间片
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}