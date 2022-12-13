package sixmode

//外观模式

type SubSystemA struct {
}

func (a SubSystemA) MethondA() {
	println("MethondA")
}

type SubSystemB struct {
}

func (b SubSystemB) MethondB() {
	println("MethondB")
}

type SubSystemC struct {
}

func (c SubSystemC) MethondC() {
	println("MethondC")
}

type Facade struct {
	A *SubSystemA
	B *SubSystemB
	C *SubSystemC
}

func (f Facade) MethondAC() {
	f.A.MethondA()
	f.C.MethondC()
}

func (f Facade) MethondAB() {
	f.A.MethondA()
	f.B.MethondB()
}
