package main

import "log"

type InterfaceA interface {
	A1func()
	A2func()
}

type InterfaceB interface {
	Bunc()
}

type InterfaceC interface {
	Cfunc()
}

type  A struct {

}

type B struct {

}

type C struct {

}

func (a *A) A1func(){
	log.Println("type A interface A1 func")
}

func (a *A) A2func(){
	log.Println("type A interface A2 func")
}

func (a *A) A3func(){
	log.Println("type A interface A3 func")
}

func main(){
	var ia InterfaceA
	ia=&A{}

	ia.A1func()

}
