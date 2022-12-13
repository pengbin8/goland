package sixmode

type Applefactory struct {
}
type Bananafactory struct {
}
type Pearfactory struct {
}

func (a *Applefactory) CreateFruit() Fruit {
	fruit := new(Apple)
	return fruit
}
func (b *Bananafactory) CreateFruit() Fruit {
	fruit := new(Banana)
	return fruit
}
func (p *Pearfactory) CreateFruit() Fruit {
	fruit := new(Pear)
	return fruit
}
