package sixmode

import "fmt"

//观察者模式

type Listener interface {
	OnTeacherComming() //观察者得到通知后要触发的动作
}

type Notifier interface {
	AddListener(listener Listener)
	RemoverListener(listener Listener)
	Notify()
}

type Student1 struct {
	Badthing string
}

func (s *Student1) OnTeacherComming() {
	fmt.Println("Student1认真听课...")
}

func (s *Student1) DoBadthing() {
	fmt.Println("做一些坏事，例如：", s.Badthing)
}

type Student2 struct {
	Badthing string
}

func (s *Student2) OnTeacherComming() {
	fmt.Println("Student2认真听课...")
}
func (s *Student2) DoBadthing() {
	fmt.Println("做一些坏事，例如：", s.Badthing)
}

type Student3 struct {
	Badthing string
}

func (s *Student3) OnTeacherComming() {
	fmt.Println("Student3认真听课...")
}
func (s *Student3) DoBadthing() {
	fmt.Println("做一些坏事，例如：", s.Badthing)
}

type ClassMonitor struct {
	listeners []Listener
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.listeners = append(c.listeners, listener)
}

func (c *ClassMonitor) RemoverListener(listener Listener) {
	for index, i := range c.listeners {
		if i == listener {
			c.listeners = append(c.listeners[:index], c.listeners[index+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) Notify() {
	for _, cm := range c.listeners {
		cm.OnTeacherComming()
	}
}

func ExecObs() {
	student1 := Student1{Badthing: "student1"}
	student2 := Student2{Badthing: "student2"}
	student3 := Student3{Badthing: "student3"}
	fmt.Println("老师没来，随便玩")
	student1.DoBadthing()
	student2.DoBadthing()
	student3.DoBadthing()
	fmt.Println("上课了，班长给了一个眼色")
	cm := new(ClassMonitor)
	cm.AddListener(&student1)
	cm.AddListener(&student2)
	cm.AddListener(&student3)
	fmt.Println("班长给了student2一个眼色，但是他不中用啊，随后班长不管他了，把他移除通知队列")
	cm.RemoverListener(&student2)
	cm.Notify()

}
