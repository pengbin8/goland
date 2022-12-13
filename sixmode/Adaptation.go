package sixmode

import "fmt"

// 装饰者模式

type Vter interface {
	//接入电压
	UserVter()
}

// 业务类 依赖 v5接口
type Tphone struct {
	v Vter
}

func NewTphone(v Vter) *Tphone {
	return &Tphone{
		v: v,
	}
}

func (t *Tphone) Charge() {
	fmt.Println("tphone进行充电...")
	t.v.UserVter()
}

// V220 ===========  搭载220v电源的业务代码=========================================
// 被适配的角色  适配者
type V220 struct {
}

func (v V220) User220V() {
	fmt.Println("使用220v电压")
}

// 电源适配器
type Adapter220V struct {
	v220 *V220
}

func (a *Adapter220V) UserVter() {
	fmt.Println("使用适配器进行充电")
	a.v220.User220V()
}

func NewAdapterV220(v *V220) *Adapter220V {
	return &Adapter220V{
		v,
	}
}

// V5 ===========  搭载5v电源的业务代码=========================================

type V5 struct {
}

func (v V5) User5V() {
	fmt.Println("使用5v电压")
}

// 电源适配器
type Adapter5V struct {
	v5 *V5
}

func (a *Adapter5V) UserVter() {
	fmt.Println("使用适配器进行充电")
	a.v5.User5V()
}

func NewAdapterV5(v *V5) *Adapter5V {
	return &Adapter5V{
		v,
	}
}
