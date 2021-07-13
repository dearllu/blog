package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	phone string
}

type Employee struct {
	Human
	school string
	loan float32
}
//Human结构体的方法
func (h Human)SayHi()  {
	fmt.Printf("hi .i an %s you can call me %s\n",h.name,h.phone)
}
//顾客结构体的方法
func (e Employee)SayHi () {

}