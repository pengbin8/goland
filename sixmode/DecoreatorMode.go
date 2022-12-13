package sixmode

//装饰器模式和代理模式的使用场景不一样，
//比如IO流使用的是装饰者模式，可以层层增加功能。
//而代理模式则一般是用于增加特殊的功能，有些动态代理不支持多层嵌套。

/*
代理和装饰其实从另一个角度更容易去理解两个模式的区别：
1.代理更多的是强调对对象的访问控制，比如说，
访问A对象的查询功能时，
访问B对象的更新功能时，
访问C对象的删除功能时，都需要判断对象是否登陆，
那么我需要将判断用户是否登陆的功能抽提出来，并对A对象、B对象和C对象进行代理，
使访问它们时都需要去判断用户是否登陆，简单地说就是将某个控制访问权限应用到多个对象上；
2.而装饰器更多的强调给对象加强功能，
比如说要给只会唱歌的A对象添加跳舞功能，添加说唱功能等，
简单地说就是将多个功能附加在一个对象上。
*/

import "fmt"

// Phone =============抽象层===================
type Phone interface {
	Show()
}

type Decorator struct {
	phone Phone
}

// Huawei ===============实现层=================
type Huawei struct {
}

func (h *Huawei) Show() {
	fmt.Println("秀出华为手机")
}

type MoDecorator struct {
	Decorator //继承基础装饰类
}

func (m *MoDecorator) Show() {
	m.phone.Show() //调用被装饰构建的原方法
	fmt.Println("增加一个贴膜")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{
		Decorator{
			phone: phone,
		},
	}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) Show() {
	k.phone.Show()
	fmt.Println("增加一个手机壳")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{
		Decorator{
			phone: phone,
		},
	}
}
