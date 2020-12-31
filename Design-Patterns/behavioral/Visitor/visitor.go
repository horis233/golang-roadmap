package Visitor

import "fmt"

type Visitor interface {
	Visit()
}
type FBVisitor struct {

}

func (v FBVisitor) Visit(){
	fmt.Println("Visit Facebook")
}

type TiktokVisitor struct {

}

func (v TiktokVisitor) Visit(){
	fmt.Println("Visit Tiktok")
}

type ElementInterface interface {
	Accept(visitor Visitor)
}

type Element struct {

}

func (e Element) Accept(v Visitor){
	v.Visit()
}
