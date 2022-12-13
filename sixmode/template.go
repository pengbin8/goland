package sixmode

//模板方法模式

type Beverage interface {
	BoilWater()          //煮开水
	Brew()               //冲泡
	PourInCup()          //倒入杯中
	AddThings()          //添加辅料
	WantAddThings() bool //是否加入辅料
}

type Template struct {
	b Beverage
}

func (t *Template) NewBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

type MakeCaffee struct {
	Template
}

func (mf *MakeCaffee) BoilWater() {
	//TODO implement me
	println("mf BoilWater me")
}

func (mf *MakeCaffee) Brew() {
	//TODO implement me
	println("mf Brew me")
}

func (mf *MakeCaffee) PourInCup() {
	//TODO implement me
	println("mf PourInCup me")
}

func (mf *MakeCaffee) AddThings() {
	//TODO implement me
	println("mf AddThings me")
}

func (mf *MakeCaffee) WantAddThings() bool {
	//TODO implement me
	println("mf WantAddThings me")
	return true
}

func NewCaffee() *MakeCaffee {
	matcaffe := new(MakeCaffee)
	matcaffe.b = matcaffe
	return matcaffe
}

// Tea ==================================================

type MakeTea struct {
	Template
}

func (m MakeTea) BoilWater() {
	//TODO implement me
	println("tea BoilWater me")
}

func (m MakeTea) Brew() {
	//TODO implement me
	println("tea Brew me")
}

func (m MakeTea) PourInCup() {
	//TODO implement me
	println("tea PourInCup me")
}

func (m MakeTea) AddThings() {
	//TODO implement me
	println("tea AddThings me")
}

func (m MakeTea) WantAddThings() bool {
	//TODO implement me
	println("tea WantAddThings me")
	return false
}

func NewMakeTea() *MakeTea {
	marketer := new(MakeTea)
	marketer.b = marketer
	return marketer
}
