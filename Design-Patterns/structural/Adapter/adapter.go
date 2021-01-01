package Adapter

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {

}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("Charging...")
}

type Adapter struct{
	*Adaptee
}

func (a *Adapter) Execute(){
	a.SpecificExecute()
}