package main

import (
	"fmt"
)

func fib(c,quit chan int) {
	x,y := 1,1
	for{
		select {
		case  c<-x:
			fmt.Println("向C中插入数据")
			x,y = y,x+y
			case <-quit:
				fmt.Println("结束了")
				fmt.Println("quit")
				return
		}


	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("执行了一个内存空间")
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0 //当向QUIT内存入数据的时候代表准备结束
	}()
	fib(c,quit)
}
