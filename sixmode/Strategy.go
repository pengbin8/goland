package sixmode

//策略模式
import "fmt"

type Strategy interface {
	UseWeapon()
}

type Ak47 struct {
}

func (s *Ak47) UseWeapon() {
	fmt.Println("使用Ak47")
}

type Ak48 struct {
}

func (a Ak48) UseWeapon() {
	fmt.Println("使用Ak48")
}

type Hero struct {
	s Strategy
}

func (h *Hero) SetStrategy(s Strategy) {
	h.s = s
}
func (h *Hero) Fight() {
	h.s.UseWeapon()
	fmt.Println("开火....")
}

func Exec() {
	hero := new(Hero)
	hero.SetStrategy(new(Ak47))
	hero.Fight()
	hero.SetStrategy(new(Ak48))
	hero.Fight()
}

/*
练习：
	商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
*/

type StrategyBase interface {
	Sale(price float64) float64
}

type StrategyA struct {
}

func (s *StrategyA) Sale(price float64) float64 {
	fmt.Println("策略A")
	return price * 0.8
}

type StrategyB struct {
}

func (s *StrategyB) Sale(price float64) float64 {
	fmt.Println("策略B")
	if price > 100 {
		price -= 100
	} else {
		price = price * 0.6
	}
	return price
}

type System struct {
	s StrategyBase
}

func (s *System) SetStrategyBase(b StrategyBase) {
	s.s = b
}

func (s *System) Sale(price float64) {
	priceSale := s.s.Sale(price)
	fmt.Println("开始销售", priceSale)
}

func ExecSystem(price float64) {
	base := new(System)
	base.SetStrategyBase(new(StrategyA))
	base.Sale(price)
	base.SetStrategyBase(new(StrategyB))
	base.Sale(price)
}
