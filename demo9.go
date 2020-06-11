package goDemos

import "fmt"

//测试 传值/引用
func PrintDemo9() {
	counter := 0
	printByValue(counter)
	fmt.Println(counter)
	printByPointer(&counter)
	fmt.Println(counter)
}

func printByValue(c int) {
	c++
	fmt.Println(c)
}

func printByPointer(c *int) {
	*c++
	fmt.Println(*c)
}
