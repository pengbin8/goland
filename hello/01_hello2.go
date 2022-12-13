package hello

import "fmt"

func init() { // 3
	fmt.Println("执行main包中main.go中init函数")
}
func Test(f func(s string)) {
	f("hello lnj")
}

func Sum(a int, b int) int {
	return a + b
}

func Calculate(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
