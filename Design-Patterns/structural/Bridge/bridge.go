package Bridge

import (
	"fmt"
)

type Draw interface {
	DrawCircle(radius, x, y int)
}

type RedCircle struct {
}

func (g *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

type YellowCircle struct {
}

func (g *YellowCircle) DrawCircle(radius, x, y int) {
	fmt.Println("radius、x、y:", radius, x, y)
}

type Shape struct {
	draw Draw
}

func (s *Shape) Shape(d Draw) {
	s.draw = d
}

type Circle struct {
	shape  Shape
	x      int
	y      int
	radius int
}

func (c *Circle) Constructor(x int, y int, radius int, draw Draw) {
	c.x = x
	c.y = y
	c.radius = radius
	c.shape.Shape(draw)
}

func (c *Circle) Draw() {
	c.shape.draw.DrawCircle(c.radius, c.x, c.y)
}
