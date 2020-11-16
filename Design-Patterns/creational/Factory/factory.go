package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type MacDonald struct {
}

func (m *MacDonald) GetFood() {
	fmt.Println("Your MacDonald is ready")
}

type BungerKing struct {
}

func (b *BungerKing) GetFood() {
	fmt.Println("Your BungerKing is ready")
}

func NewRestaurant(s string) Restaurant {
	switch s {
	case "m":
		return &MacDonald{}
	case "b":
		return &BungerKing{}
	}
	return nil
}
