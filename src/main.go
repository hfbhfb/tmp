package main

import (
	"fmt"
	"reflect"
)

type Ia interface {
	A()
}

type Eggo struct {
	Name string
}

func (e *Eggo) A() {
	fmt.Println("eggo A")
}
func (e *Eggo) B() {
	fmt.Println("eggo A")
}

func (e *Eggo) C() Ia {
	fmt.Println("auto change to Interface Type Ia")
	return e
}

func main() {
	a := Eggo{Name: "eggo name"}

	a.C().A()

	t := reflect.TypeOf(a)
	println(t.Name(), t.NumField())
	fmt.Println("aaa")
}
