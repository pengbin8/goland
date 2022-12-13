package sixmode

type Fruit interface {
	Show() string
}

type Apple struct {
}

func (a *Apple) Show() string {
	return "我是苹果"
}

type Banana struct {
	Name string
}

func (b *Banana) Show() string {
	b.Name = "我是香蕉"
	return b.Name
}

type Pear struct {
}

func (p *Pear) Show() string {
	return "我是李子"
}

type SimpleFactory struct {
}

func (f SimpleFactory) Createfruit(kind string) Fruit {
	var fruit Fruit
	switch kind {
	case "apple":
		fruit = new(Apple)
	case "pear":
		fruit = new(Pear)
	case "banana":
		fruit = new(Banana)

	}
	return fruit
}
