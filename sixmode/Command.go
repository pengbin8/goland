package sixmode

//命令模式

import "fmt"

type Doctor struct {
}

func (d *Doctor) treatEye() {
	fmt.Println("治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("治疗鼻子")
}

type Command interface {
	Treat()
}

type CommandTreatEye struct {
	d *Doctor
}

func (c *CommandTreatEye) Treat() {
	c.d.treatEye()
}

type CommandTreatNose struct {
	d *Doctor
}

func (c *CommandTreatNose) Treat() {
	c.d.treatNose()
}

type Nurse struct {
	CmdList []Command
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}

func BingRen() {
	docter := new(Doctor)
	cmdEye := CommandTreatEye{docter}
	cmdNose := CommandTreatNose{docter}
	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)
	nurse.Notify()

}

/*
练习：
路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员。
根据命令模式，设计烤串场景
*/

type Cooker struct {
}

func (c *Cooker) CommandCookerChicken() {
	fmt.Println("烤一只鸡")
}

func (c *Cooker) CommandCookerDuck() {
	fmt.Println("烤一只鸭子")
}

type Command2 interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (c *CommandCookChicken) Make() {
	c.cooker.CommandCookerChicken()
}

type CommandCookDuck struct {
	cooker *Cooker
}

func (c *CommandCookDuck) Make() {
	c.cooker.CommandCookerDuck()
}

type WaiterMm struct {
	CmdList []Command2
}

func (w *WaiterMm) Notify() {
	if w.CmdList == nil {
		fmt.Println("w.CmdList == nil")
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func Consumer() {
	cooker := new(Cooker)
	cmdCookerChicken := CommandCookChicken{
		cooker: cooker,
	}
	cmdCookerDuck := CommandCookDuck{
		cooker: cooker,
	}
	mm := new(WaiterMm)
	mm.CmdList = append(mm.CmdList, &cmdCookerDuck)
	mm.CmdList = append(mm.CmdList, &cmdCookerChicken)
	mm.Notify()
}
