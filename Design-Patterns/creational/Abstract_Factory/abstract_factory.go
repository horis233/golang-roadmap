package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}

type rice struct {
}

func (r *rice) Cook() {
	fmt.Println("it is a rice.")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("it is Tomato.")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type simpleLunchFactory struct {
}

func NewSimpleShapeFactory() LunchFactory {
	return &simpleLunchFactory{}
}

func (s *simpleLunchFactory) CreateFood() Lunch {
	return &rice{}
}

func (s *simpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
