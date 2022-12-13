package sixmode

type AbstractFactory interface {
	CreateFruit() Fruit
}

func NewFruit(a AbstractFactory) Fruit {
	return a.CreateFruit()
}
