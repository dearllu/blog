package main

import "runtime/metrics"

func sum(a []int, c chan int) {
	total := 0
	for _,v := range a{
		total += v
	}
	c <- total
}


func main() {
	a := []int{7,2,8,-9,0}
	c := make(chan int) //不带缓冲

}