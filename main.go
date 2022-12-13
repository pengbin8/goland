package main

import (
	"fmt"
	"github.com/rs/xid"
	"goland/hello"
	"goland/sixmode"
)

func main() {

	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id)

	// 1.定义一个数组
	arr := [3]int{1, 3, 5}
	// 2.快速遍历数组
	// i用于保存当前遍历到数组的索引
	// v用于保存当前遍历到数组的值
	for i, v := range arr {
		fmt.Println(i, v)
	}

	var sumvalue = hello.Sum(1, 2)
	fmt.Println("xxxx", sumvalue)

	var sum1, sub1 int = hello.Calculate(10, 100)
	fmt.Println("calculate", sum1, sub1)

	var person0 = hello.Person{}
	fmt.Println("person0:", person0)
	person := hello.Person{
		2.5, "test_go",
	}
	fmt.Println("person:", person.Age, person.Name)

	hello.Test(func(s string) {
		fmt.Println(s)
	})

	// 1.定义的同时完全初始化
	var arr1 = [...]int{1, 3, 5}
	// 2.打印数组
	fmt.Println(arr1) // [1 3 5]

	// 1.定义的同时指定元素初始化
	var arr2 = [...]int{6: 5, 5: 4}
	// 2.打印数组
	fmt.Println(arr2) // [0 0 0 0 0 0 5]

	var arrsce = [6]int{1, 3, 5, 7, 9}
	// 同时指定开始位置和结束位置
	var sce1 = arrsce[0:2]
	fmt.Println(sce1)                   // [1 3]
	fmt.Println("cap-sce1:", cap(sce1)) // 5

	var sce2 = arrsce[5:6]
	fmt.Println(sce2)                   // [0]
	fmt.Println("cap-sce2:", cap(sce2)) // 1

	// 第一个参数: 指定切片数据类型
	// 第二个参数: 指定切片的长度
	// 第三个参数: 指定切片的容量
	//var sce = make([]int, 3, 5)
	var sce = []int{1, 3, 5}
	fmt.Println(sce)      // [1 3 5]
	fmt.Println(len(sce)) // 3
	fmt.Println(cap(sce)) // 3

	var dict = map[string]string{"name": "lnj", "age": "33", "gender": "male"}
	//value, ok := dict["age"]
	//if(ok){
	//	fmt.Println("有age这个key,值为", value)
	//}else{
	//	fmt.Println("没有age这个key,值为", value)
	//}
	if value, ok := dict["age"]; ok {
		fmt.Println("有age这个key,值为", value)
	}

	type Demo struct {
		age int               // 基本类型作为属性
		arr [3]int            // 数组类型作为属性
		sce []int             // 切片类型作为属性
		mp  map[string]string // 字典类型作为属性
		stu hello.Student     // 结构体类型作为属性
	}

	var (
		d Demo = Demo{
			age: 33,
			arr: [3]int{1, 3, 5},
			sce: []int{2, 4, 6},
			mp:  map[string]string{"class": "one"},
			stu: hello.Student{
				"lnj",
				33,
			},
		}
	)

	fmt.Println("ddddd:", d)

	fmt.Println("::::::::::::::::::::::::::::::::::")
	fmt.Println("::::::::::::::单例::::::::::::::start")
	var instance1 = sixmode.Getinstance()
	fmt.Println("instance :", *instance1)
	var instance2 = sixmode.Getinstance()
	fmt.Println("instance :", *instance2)
	fmt.Println("::::::::::::::单例::::::::::::::end")

	fmt.Println("::::::::::::::简单工厂::::::::::::::start")
	var factory = sixmode.SimpleFactory{}
	var showname = factory.Createfruit("pear").Show()
	fmt.Println("SimpleFactory :", showname)
	fmt.Println("::::::::::::::简单工厂::::::::::::::end")

	fmt.Println("::::::::::::::普通工厂::::::::::::::start")
	app := sixmode.Applefactory{}
	fmt.Println("Applefactory :", app.CreateFruit().Show())
	fmt.Println("::::::::::::::普通工厂::::::::::::::end")

	fmt.Println("::::::::::::::抽象工厂::::::::::::::start")
	var abstractfactory sixmode.AbstractFactory
	abstractfactory = new(sixmode.Pearfactory)
	abspear := sixmode.NewFruit(abstractfactory).Show()
	fmt.Println("AbstractFactory：", abspear)
	fmt.Println("::::::::::::::抽象工厂::::::::::::::end")

	fmt.Println("::::::::::::::代理模式::::::::::::::start")
	koreaShopping := new(sixmode.AmericanShoping)
	proxy := sixmode.NewProxy(koreaShopping)
	buy := sixmode.Goods{
		"包包",
		true,
	}
	proxy.Buy(&buy)

	fmt.Println("::::::::::::::代理模式::::::::::::::end")

	fmt.Println("::::::::::::::装饰者模式::::::::::::::start")

	var huawei sixmode.Phone
	huawei = new(sixmode.Huawei)
	dohuawei := sixmode.NewMoDecorator(huawei)
	kehuawei := sixmode.NewKeDecorator(dohuawei)
	kehuawei.Show()
	fmt.Println("::::::::::::::装饰者模式::::::::::::::end")

	fmt.Println("::::::::::::::适配器模式::::::::::::::start")
	v220 := new(sixmode.V220)
	zhuangshizhemode := sixmode.NewAdapterV220(v220)
	sixmode.NewTphone(zhuangshizhemode).Charge()
	fmt.Println("----------------------------------")
	v5 := new(sixmode.V5)
	zhuangshizhemodeV5 := sixmode.NewAdapterV5(v5)
	sixmode.NewTphone(zhuangshizhemodeV5).Charge()

	fmt.Println("::::::::::::::适配器模式::::::::::::::end")

	fmt.Println("::::::::::::::外观模式::::::::::::::start")

	f := sixmode.Facade{
		A: new(sixmode.SubSystemA),
		B: new(sixmode.SubSystemB),
		C: new(sixmode.SubSystemC),
	}
	f.MethondAB()
	fmt.Println("---------------------------------")
	f.MethondAC()

	fmt.Println("::::::::::::::外观模式::::::::::::::end")

	mf := sixmode.NewCaffee()
	mf.NewBeverage()
	fmt.Println("---------------------------------")
	mt := sixmode.NewMakeTea()
	mt.NewBeverage()

	fmt.Println("---------------------------------")
	sixmode.BingRen()

	fmt.Println("-----------------命令模式----------------")
	sixmode.Consumer()
	fmt.Println("-----------------命令模式----------------")

	fmt.Println("-----------------策略模式----------------")
	sixmode.Exec()
	fmt.Println("--------")
	sixmode.ExecSystem(58)
	fmt.Println("-----------------策略模式----------------")

	sixmode.ExecObs()

	fmt.Println("****************************************SIX MODE******************************")

}
