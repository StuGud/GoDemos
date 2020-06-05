package goDemos

import "fmt"

//测试双引号 反引号 以及二者拼接对转义符的影响

func PrintDemo1() {
	str1 := "abc"
	str2 := `abc
efg
			\n hij`
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str1 + str2)
	fmt.Println(str1 + "\n" + str2)
}
